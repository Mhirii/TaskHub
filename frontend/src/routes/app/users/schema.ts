import { z } from "zod";

export const userSchema = z.object({
	username: z.string().min(1).max(255),
	email: z.string().email().min(1).max(255),
	password: z.string().min(1).max(255),
	roles: z.array(z.string()).min(1),
})

export type UserSchema = typeof userSchema;
