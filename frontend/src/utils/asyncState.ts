export function createAsyncState<T>(initialData: T) {
  return {
    data: initialData,
    loading: false,
    error: null as string | null
  }
}