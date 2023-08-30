<script lang="ts">
	import Icon from '@krowten/svelte-heroicons/Icon.svelte';
import { getContext, onDestroy } from 'svelte';
	import type { Writable } from 'svelte/store';
	import { fade } from 'svelte/transition';
    import type {CssClasses} from "./stypes"
	// Types

	// Props
	export let locked: boolean = false;
	export let back_locked: boolean = false;

	// Props (regions)
	/** Provide arbitrary classes to the step header region. */
	export let regionHeader: CssClasses = '';
	/** Provide arbitrary classes to the step content region. */
	export let regionContent: CssClasses = '';
	/** Provide arbitrary classes to the step navigation region. */
	export let regionNavigation: CssClasses = '';
	// Context
	export let state: Writable<any> = getContext('state');
	export let dispatchParent: any = getContext('dispatchParent');
	export let stepTerm: string = getContext('stepTerm');
	export let gap: CssClasses = getContext('gap');
	export let justify: CssClasses = getContext('justify');
	export let buttonBack: CssClasses = getContext('buttonBack');
	export let buttonBackType: 'submit' | 'reset' | 'button' = getContext('buttonBackType');
	export let buttonBackLabel: string = getContext('buttonBackLabel');
	export let buttonNext: CssClasses = getContext('buttonNext');
	export let buttonNextType: 'submit' | 'reset' | 'button' = getContext('buttonNextType');
	export let buttonNextLabel: string = getContext('buttonNextLabel');
	export let buttonComplete: CssClasses = getContext('buttonComplete');
	export let buttonCompleteType: 'submit' | 'reset' | 'button' = getContext('buttonCompleteType');
	export let buttonCompleteLabel: string = getContext('buttonCompleteLabel');
	// Register step on init (keep these paired)
	const stepIndex = $state.total;
	$state.total++;
	// Classes
	const cBase = 'space-y-4';
	const cHeader = 'text-2xl font-bold';
	const cContent = 'space-y-4';
	const cNavigation = 'flex';
	function onNext(): void {
		if (locked) return;
		$state.current++;
		/** @event { $state } next - Fires when the NEXT button is pressed per step.  */
		dispatchParent('next', { step: stepIndex, state: $state });
		dispatchParent('step', { step: stepIndex, state: $state });
	}
	function onBack(): void {
		if (back_locked) return
		$state.current--;
		/** @event { $state } back - Fires when the BACK button is pressed per step.  */
		dispatchParent('back', { step: stepIndex, state: $state });
		dispatchParent('step', { step: stepIndex, state: $state });
	}
	function onComplete() {
		/** @event { $state } complete - Fires when the COMPLETE button is pressed.  */
		dispatchParent('complete', { step: stepIndex, state: $state });
	}
	// Reactive
	$: classesBase = `${cBase} ${$$props.class ?? ''}`;
	$: classesHeader = `${cHeader} ${regionHeader}`;
	$: classesContent = `${cContent} ${regionContent}`;
	$: classesNavigation = `${cNavigation} ${justify} ${gap} ${regionNavigation}`;
	// Unregister step on destroy
	onDestroy(() => {
		$state.total--;
	});
</script>
{#if stepIndex === $state.current}
	<div class="step {classesBase}" data-testid="step">
		<!-- Slot: Header -->
		<header class="step-header {classesHeader}">
			<slot name="header">{stepTerm} {stepIndex + 1}</slot>
		</header>
		<!-- Slot: Default -->
		<div class="step-content {classesContent}">
			<slot>({stepTerm} {stepIndex + 1} Content)</slot>
		</div>
		<!-- Navigation -->
		{#if $state.total > 1}
			<div class="step-navigation {classesNavigation}" transition:fade|local={{ duration: 100 }}>
				<button type={buttonBackType} class="btn inline-flex bg-gray-600 hover:bg-gray-800 rounded text-white {buttonBack}" on:click={onBack} disabled={$state.current === 0}>
					{#if back_locked}
						<Icon name="lock-closed" class="w-3 aspect-square" />
					{/if}

					{@html buttonBackLabel}
				</button>
				{#if stepIndex < $state.total - 1}
					<button type={buttonNextType} class="btn inline-flex {buttonNext} bg-gray-600 hover:bg-gray-800 rounded-full p-2 text-white" on:click={onNext} disabled={locked}>
						{#if locked}
							<Icon name="lock-closed" class="w-3 aspect-square" />
						{/if}
						<span>{@html buttonNextLabel}</span>
					</button>
				{:else}
					<button type={buttonCompleteType} class="btn inline-flex {buttonComplete}" on:click={onComplete}>
						{#if locked}
							<Icon name="lock-closed" class="w-3 aspect-square" />
						{/if}
						<span>{@html buttonCompleteLabel}</span>
					</button>
				{/if}
			</div>
		{/if}
	</div>
{/if}

