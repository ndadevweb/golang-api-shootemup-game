import type { SigninFormInterface } from "../types/types";

export async function signin(payload: SigninFormInterface) {
  const response = await fetch("/api/auth/signin", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  })

  if (!response.ok) {
    const { error } = await response.json()

    throw Error(error)
  }

  return await response.json()
}
