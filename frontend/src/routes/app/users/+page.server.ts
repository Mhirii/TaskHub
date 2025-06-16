import { USERS_API_URL } from "$env/static/private";
import type { GetUsersResponse } from "$lib/types";
import { superValidate } from "sveltekit-superforms";
import { userSchema } from "./schema";
import { zod } from "sveltekit-superforms/adapters";
import { fail, redirect } from "@sveltejs/kit";

export async function load({ cookies }) {
	const access_token = cookies.get('access_token');
	const getUsers = async () => fetch(`${USERS_API_URL}/`, {
		method: 'GET',
		headers: {
			'Content-Type': 'application/json',
			'Authorization': `Bearer ${access_token}`,
		}
	}).then(res => {
		if (!res.ok) throw new Error('Failed to fetch users');
		const result = res.json()
		console.log(result)
		return result
	});
	const userForm = await superValidate(zod(userSchema))
	const users: GetUsersResponse = await getUsers()
	return {
		users,
		userForm,
	};
}

export const actions = {
	default: async ({ cookies, request }) => {
		const formData = await request.formData();
		const formType = formData.get("formType")?.toString();

		if (!formType) {
			return fail(400, { message: "Form type is required" });
		}
		const access_token = cookies.get('access_token');
		if (!access_token) {
			redirect(302, '/auth/login');
		}

		switch (formType) {
			case "createUser": {
				const form = await superValidate(formData, zod(userSchema));
				if (!form.valid) {
					return fail(400, {
						form,
					});
				}
				const { username, email, password, roles } = form.data;
				const body = JSON.stringify({ username, email, password, roles })
				const res = await fetch(`${USERS_API_URL}/`,
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
					userId: resp.id
				};
			}
		}
	}
};
