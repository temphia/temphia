import { writable, Writable } from "svelte/store";
import type { ActionResponse, Environment } from "../../../lib";
import { StageStore } from "./stage";
import type { State, Manager } from "./wizard_types";

export class WizardManager implements Manager {
  wizard_title?: string;
  _env: Environment;
  _opaqueData?: string;
  _state: Writable<State>;
  _stage_Store: StageStore;
  _exec_options?: any;

  constructor(env: Environment, opts?: any) {
    this._env = env;
    this._state = writable({
      data_sources: {},
      fields: [],
      final: false,
      flowState: "NOT_LOADED",
      epoch: 0,
    });

    this._exec_options = opts;
    this._state.subscribe((state) => console.log("STATE =>", state));
    this._stage_Store = null;
  }

  get_state = () => {
    return this._state;
  };

  get_field_store = (field: string) => {
    return this._stage_Store.get_field_store(field);
  };

  init = async () => {
    const resp = await this._env.PreformAction("get_splash", {
      has_exec_data: !!this._exec_options,
    });

    if (!resp.status_ok) {
      console.warn("error getting wizard splash", resp);
      return;
    }

    console.log("INIT RESP", resp);
    this.applySplashFields(resp);
    if (resp.body["skip_splash"]) {
      this.splash_next();
    }
  };

  applySplashFields = (resp: ActionResponse) => {
    this._stage_Store = new StageStore(this);

    this.wizard_title = resp.body["wizard_title"] || "";
    const fields = resp.body["fields"] || [];
    const message = resp.body["message"] || "";
    const data_sources = resp.body["data_sources"] || {};
    this._state.update((old) => ({
      ...old,
      fields,
      epoch: old.epoch + 1,
      message,
      data_sources,
      flowState: "SPLASH_LOADED",
    }));
  };

  splash_next = async () => {
    const values = this._stage_Store._get_values();

    const resp = await this._env.PreformAction("run_start", {
      splash_data: values,
      start_raw_data: this._exec_options,
    });
    if (!resp.status_ok) {
      console.warn("error starting from splash", resp);
      return;
    }

    if (resp.body["stage_started"]) {
      this.applyStageFields(resp);
    } else {
      this.applySplashFields(resp);
    }
  };

  applyStageFields = (resp: ActionResponse) => {
    this._stage_Store = new StageStore(this);
    const fields = resp.body["fields"] || [];
    this._opaqueData = resp.body["odata"] || "";
    const stageTitle = resp.body["stage_title"];
    const data_sources = resp.body["data_sources"] || {};
    const errors = resp.body["errors"] || {};
    const prev_data = resp.body["prev_data"];

    this._state.update((old) => ({
      ...old,
      data_sources,
      epoch: old.epoch + 1,
      fields,
      stageTitle,
      flowState: "STAGE_LOADED",
      errors,
      prev_data,
    }));
  };

  stage_next = async () => {
    const values = this._stage_Store._get_values();
    this._state.update((old) => ({ ...old, flowState: "STAGE_PROCESSING" }));

    const resp = await this._env.PreformAction("run_next", {
      data: values,
      odata: this._opaqueData,
    });
    if (!resp.status_ok) {
      console.warn("error going to next stage", resp);
      return;
    }

    console.log("@=>", resp);

    if (!resp.body["ok"]) {
      const errors = resp.body["errors"] || {};
      this._state.update((old) => ({
        ...old,
        errors,
        prev_data: values,
        flowState: "STAGE_LOADED",
      }));
      return;
    }

    if (resp.body["final"]) {
      const message = resp.body["last_message"] || "";
      this._state.update((old) => ({ ...old, flowState: "FINISHED", message }));
      return;
    }

    this.applyStageFields(resp);
  };

  stage_back = async () => {

  };
}
