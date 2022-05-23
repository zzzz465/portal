import ono from 'ono'
import { Option } from '../types/option'

type Func<T> = () => T

export function try0<E = {}, T = {}>(func: Func<T>): Option<T, E> {
  try {
    return [func(), null]
  } catch (err) {
    return [null, ono(err as E)]
  }
}

/**
 * wraps method with try/catch. catches the error and returns as Option<T, E>
 * an error may be wrapped using ono if ErrorLike is being thrown.
 * otherwise, it'll return an thrown value as it is.
 */
export async function tryP<E = {}, T = any>(func: Func<T>): Promise<Option<Awaited<T>, E>> {
  try {
    return [await func(), null]
  } catch (err) {
    if (err === null || err === undefined) {
      return [null, {} as any]
    } else if (err instanceof Error) {
      return [null, ono(err as any)]
    } else {
      return [null, err as E]
    }
  }
}
