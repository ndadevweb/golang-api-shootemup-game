import type { SignupFormInterface } from "../types/types";

export async function signup(payload: SignupFormInterface) {
  const response = await fetch("/api/auth/signup", {
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
