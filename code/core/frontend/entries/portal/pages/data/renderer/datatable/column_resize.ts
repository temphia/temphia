import { get, writable } from 'svelte/store'

export const ColumnResize = (DEFAULT_WIDTH) => {
  const MIN_WIDTH = 10

  const widths = writable({})

  const setWidth = (colid, width) => {
    if (width < MIN_WIDTH) {
      width = MIN_WIDTH
    }

    widths.update((old) => {
      return { ...old, [colid]: width }
    })
  }

  const getHandler = (colid) => {
    return (ev) => {
      let basePageX = ev.pageX

      const mouseMove = (ev) => {
        if (basePageX === 0) return

        let width = get(widths)[colid] || DEFAULT_WIDTH
        width = width + (ev.pageX - basePageX) * 0.003
        setWidth(colid, width)
      }

      const mouseUp = () => {
        basePageX = 0
        document.removeEventListener('mouseup', mouseUp)
        document.removeEventListener('mousemove', mouseMove)
      }

      document.addEventListener('mousemove', mouseMove)
      document.addEventListener('mouseup', mouseUp)
    }
  }

  return { subscribe: widths.subscribe, getHandler }
}
