import type { PageServerLoad } from "./$types";
import { PROJECTS_API_URL } from "$env/static/private";
import type { GetProjectsResponse, Project } from "$lib/types";

export const load: PageServerLoad = async ({ cookies, fetch }) => {
	const access_token = cookies.get('access_token');
	const projectsRes: Promise<GetProjectsResponse> = fetch(`${PROJECTS_API_URL}/`, {
		method: 'GET',
		headers: {
			'Authorization': `Bearer ${access_token}`,
			'Content-Type': 'application/json'
		}
	}).then(res => {
		if (!res.ok) throw new Error('Failed to fetch projects');
		return res.json();
	});
	const projects: Promise<Project[]> = projectsRes.then((p) => p?.map((p) => ({
		...p,
		created_at: new Date(p.created_at * 1000).toISOString(),
		updated_at: new Date(p.updated_at * 1000).toISOString(),
	}) || []));
	return { projects };
};

