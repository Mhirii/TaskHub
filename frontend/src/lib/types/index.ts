
export type Project = {
	id: number;
	name: string;
	description: string;
	ownerID: number;

	task_count: number;
	open_task_count: number;
	members_count: number;

	created_at: string;
	updated_at: string;
}
export type GetProjectsResponse = Array<Omit<Project, "created_at" | "updated_at"> & { created_at: number, updated_at: number }>;

export type Task = {
	id: number;
	name: string;
	description: string;

	board_id: number;
	project_id: number;

	order: number;
	parent_id: number;
	status: number;
	priority: number;
	due_date: number;
	assignee_id: number;
	tags: string[];

	created_by: number;
	created_at: number;
	updated_at: number;
}
