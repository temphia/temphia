package runtime

func (r *runtime) controlLoop() {

	for {

		select {
		case <-r.close:
			return
		case j := <-r.jobCh:

			ns := r.getNS(j.Inner.Namespace, true)

			resp, err := ns.doWork(&j.Inner)
			if err != nil {
				j.Err(err)
				return
			}

			j.Ok(resp)
		}

	}

}
