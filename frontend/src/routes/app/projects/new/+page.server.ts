import { superValidate } from "sveltekit-superforms";
import type { Actions, PageServerLoad } from "./$types";
import { zod } from "sveltekit-superforms/adapters";
import { projectSchema } from "./schema";
import { fail, redirect } from "@sveltejs/kit";
import { PROJECTS_API_URL } from "$env/static/private";

export const load: PageServerLoad = async () => {
	return {
		form: await superValidate(zod(projectSchema)),
	};
};

export const actions: Actions = {
	default: async ({ cookies, request, fetch }) => {
		const access_token = cookies.get('access_token');
		if (!access_token) {
			redirect(302, '/auth/login');
		}
		const formData = await request.formData();
		const form = await superValidate(formData, zod(projectSchema));
		if (!form.valid) {
			return fail(400, {
				form,
			});
		}

		const { name, description } = form.data;
		const body = JSON.stringify({
			name,
			description,
		})

		const res = await fetch(`${PROJECTS_API_URL}/`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				'Authorization': `Bearer ${access_token}`,
			},
			body,
		});
		if (!res.ok) return fail(res.status, { form });
		const resp: { id: number } = await res.json();
		redirect(302, `/app/projects/${resp.id}`);
	},
};
