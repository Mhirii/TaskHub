import { redirect } from "@sveltejs/kit";

export function load({ cookies }) {
	const access_token = cookies.get('access_token');
	redirect(302, '/app/projects');


	return {
		access_token,
	};
}

