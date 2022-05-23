import ono, { OnoError } from 'ono'
import { Option } from '../types/option'

type Func<T> = () => T

export function try0<T, E>(func: Func<T>): Option<T, E> {
  try {
    return [func(), null]
  } catch (err) {
    return [null, ono(err as E)]
  }
}

export async function tryP<T, E>(func: Func<T>): Promise<Option<Awaited<T>, OnoError<E>>> {
  try {
    return [await func(), null]
  } catch (err) {
    // TODO: maybe err is not an error type
    return [null, ono(err as E)]
  }
}
