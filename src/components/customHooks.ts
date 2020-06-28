import { useEffect, useRef } from 'react'
export const useDidMountEffect = (fn: any, deps: any) => {
  const didMountRef = useRef(false)
  useEffect(() => {
    if (didMountRef.current) fn()
    else didMountRef.current = true
  }, deps)
}
