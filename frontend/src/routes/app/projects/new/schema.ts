import { z } from "zod";


export const projectSchema = z.object({
	name: z.string().min(1).max(255),
	description: z.string().min(1).max(255),
})


export type ProjectSchema = typeof projectSchema;
