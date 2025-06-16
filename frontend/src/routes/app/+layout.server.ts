import { PROJECTS_API_URL } from "$env/static/private";
import type { GetProjectsResponse, Project } from "$lib/types";
import type { LayoutServerLoad, } from "./$types";
import { redirect } from '@sveltejs/kit'
import { decode } from "jsonwebtoken"


type JWTPayload = {
	exp: number;
	iat: number;
	iss: string;
	sub: string;
	roles: string[];
	username: string;
	[key: string]: any;
}
export const load: LayoutServerLoad = async ({ cookies }) => {
	const access_token = cookies.get('access_token');
	if (!access_token) {
		redirect(302, '/auth/login');
	}

	const j: JWTPayload | null | string = decode(access_token);
	if (!j) {
		redirect(302, '/auth/login');
	}
	if (typeof j === "string") {
		console.log("JWT: ", access_token, " is invalid : ", j)
		redirect(302, '/auth/login');
	}
	console.log(j);

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
	return { projects, user: { username: j.username, email: j.sub, avatar: "", roles: j.roles } };
};
