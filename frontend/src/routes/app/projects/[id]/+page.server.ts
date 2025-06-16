
import type { PageServerLoad } from "./$types";
import { PROJECTS_API_URL, TASKS_API_URL } from "$env/static/private";
import type { Task } from "$lib/types";
import { fail, redirect, type Actions, type RequestEvent } from "@sveltejs/kit";
import type { GetProjectsResponse, Project } from "$lib/types";
import { superValidate } from "sveltekit-superforms";
import { zod, type ZodValidation } from "sveltekit-superforms/adapters";
import { taskSchema, taskUpdateShema } from "./schema";


export const load: PageServerLoad = async ({ cookies, fetch, params }) => {
	const projectID = params.id;
	const access_token = cookies.get('access_token');

	const headers = {
		'Authorization': `Bearer ${access_token}`,
		'Content-Type': 'application/json'
	}

	const tasks: Promise<Task[]> = fetch(`${TASKS_API_URL}/project/${projectID}`, {
		headers,
	}).then(res => {
		if (!res.ok) fail(res.status, { message: `Failed to fetch tasks for project ${projectID}: ${res.statusText}` });
		return res.json();
	})
	const project: Promise<Project> = fetch(`${PROJECTS_API_URL}/${projectID}`, {
		headers,
	}).then(async (res) => {
		if (!res.ok) fail(res.status, { message: `Failed to fetch project ${projectID}: ${res.statusText}` });
		const p: Promise<GetProjectsResponse[0]> = res.json();
		return p.then((p) => ({ ...p, created_at: "", updated_at: "" }));
	});

	const taskForm = await superValidate(zod(taskSchema))
	const taskUpdateForm = await superValidate(zod(taskUpdateShema))
	return { tasks, project, taskForm, taskUpdateForm };
};


export const actions: Actions = {
	default: async (event) => {
		const formData = await event.request.formData();
		const formType = formData.get("formType")?.toString();

		if (!formType) {
			return fail(400, { message: "Form type is required" });
		}
		const access_token = event.cookies.get('access_token');
		if (!access_token) {
			redirect(302, '/auth/login');
		}

		switch (formType) {
			case "createTask": {
				const form = await superValidate(formData, zod(taskSchema));
				if (!form.valid) {
					return fail(400, {
						form,
					});
				}
				const { name, description, due_date, priority, assignee_id, board_id, project_id } = form.data;
				const body = JSON.stringify({
					name,
					description,
					due_date: due_date ? Math.ceil(due_date.valueOf() / 1000) : null,
					priority,
					assignee_id,
					board_id,
					project_id,
				})
				const res = await fetch(`${TASKS_API_URL}/`,
					{
						method: 'POST',
						headers: {
							'Content-Type': 'application/json',
							'Authorization': `Bearer ${access_token}`,
						},
						body,
					});
				if (!res.ok) return fail(res.status, { form });
				const resp: { id: number } = await res.json();
				return {
					form,
					taskId: resp.id
				};
			}
			case "updateTask": {
				console.log(formData)
				const form = await superValidate(formData, zod(taskUpdateShema));
				if (!form.valid) {
					return fail(400, {
						form,
					});
				}
				const { id, name, description, due_date, priority, assignee_id, board_id, project_id, status } = form.data;
				const body = JSON.stringify({
					id,
					name,
					description,
					due_date: due_date ? Math.ceil(due_date.valueOf() / 1000) : null,
					priority,
					assignee_id,
					board_id,
					project_id,
					status,
				})
				console.log("using ", body)
				const res = await fetch(`${TASKS_API_URL}/${id}`,
					{
						method: 'PATCH',
						headers: {
							'Content-Type': 'application/json',
							'Authorization': `Bearer ${access_token}`,
						},
						body,
					});
				if (!res.ok) return fail(res.status, { form });
				const resp: { id: number } = await res.json();
				return {
					form,
					taskId: resp.id
				};
			}
		}
	}
};

