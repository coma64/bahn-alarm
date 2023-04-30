// Impossible to type this better as the angular compiler will convert the looped over type
// to `Original & { id: unknown }`
// eslint-disable-next-line @typescript-eslint/no-unsafe-return
export const trackById = (_: number, { id }: any) => id;
