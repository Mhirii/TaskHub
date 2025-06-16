import { z } from "zod";

export const taskSchema = z.object({
	name: z.string().min(1),
	description: z.string().min(1),
	due_date: z.optional(z.date()),
	priority: z.number(),
	assignee_id: z.number(),
	board_id: z.number(),
	project_id: z.number(),
	status: z.number(),
});
export type TaskSchema = typeof taskSchema;

export const taskUpdateShema = z.object({
	id: z.number(),
	name: z.string().min(1),
	description: z.string().min(1),
	due_date: z.optional(z.date()),
	priority: z.optional(z.number()),
	assignee_id: z.nullable(z.optional(z.number())),
	board_id: z.number(),
	project_id: z.number(),
	status: z.optional(z.number()),
});

export type TaskUpdateSchema = typeof taskUpdateShema;
