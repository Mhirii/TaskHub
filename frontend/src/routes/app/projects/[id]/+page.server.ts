
import type { PageServerLoad } from "./$types";
import { PROJECTS_API_URL, TASKS_API_URL } from "$env/static/private";
import type { Task } from "$lib/types";
import { fail } from "@sveltejs/kit";
import type { GetProjectsResponse, Project } from "$lib/types";


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
	});
	const project: Promise<Project> = fetch(`${PROJECTS_API_URL}/${projectID}`, {
		headers,
	}).then(async (res) => {
		if (!res.ok) fail(res.status, { message: `Failed to fetch project ${projectID}: ${res.statusText}` });
		const p: Promise<GetProjectsResponse[0]> = res.json();
		return p.then((p) => ({ ...p, created_at: "", updated_at: "" }));
	});
	return { tasks, project };
};

