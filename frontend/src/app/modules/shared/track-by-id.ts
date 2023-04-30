// Impossible to type this better as the angular compiler will convert the looped over type
// to `Original & { id: unknown }`
export const trackById = (_: number, { id }: any) => id;
