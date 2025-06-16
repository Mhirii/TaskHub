
import type { PageServerLoad, Actions } from "./$types.js";
import { fail } from "@sveltejs/kit";
import { formSchema } from "./schema";
import { zod } from "sveltekit-superforms/adapters";
import { superValidate } from "sveltekit-superforms";


import { AUTH_API_URL } from '$env/static/private';

type LoginResponseSchema = {
	access_token: string;
	refresh_token: string;
	expires_in: number;
	user: {
		id: number,
		username: string,
		email: string,
		created_at: string,
		updated_at: string,
	}
};

export const load: PageServerLoad = async () => {
	return {
		form: await superValidate(zod(formSchema)),
	};
};

export const actions: Actions = {
	default: async ({ cookies, request }) => {
		const formData = await request.formData();
		console.log(formData)
		const form = await superValidate(formData, zod(formSchema));
		if (!form.valid) {
			return fail(400, {
				form,
			});
		}
		console.log(form.data)
		const { usernameOrEmail, password } = form.data;
		const body = JSON.stringify({
			username: usernameOrEmail,
			password,
		})
		const res = await fetch(`${AUTH_API_URL}/login`,
			{
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
				},
				body,
			});
		if (!res.ok) {
			return fail(res.status, { form });
		}
		const resp: LoginResponseSchema = await res.json();
		cookies.set('access_token', resp.access_token, { path: '/' });
		return {
			form,
			...resp,
		};
	},
};

