<script lang="ts">
	import * as Form from "$lib/components/ui/form/index.js";
	import { Input } from "$lib/components/ui/input/index.js";
	import { formSchema, type FormSchema } from "./schema";
	import {
		type SuperValidated,
		type Infer,
		superForm,
	} from "sveltekit-superforms";
	import { zodClient } from "sveltekit-superforms/adapters";
	import { setAccess, setRefresh } from "./store";
	import { goto } from "$app/navigation";
	import { userStore } from "$lib/store/user";

	let { data }: { data: { form: SuperValidated<Infer<FormSchema>> } } =
		$props();

	const form = superForm(data.form, {
		validators: zodClient(formSchema),
		onResult: async ({ result }) => {
			//@ts-ignore
			const access = result.data.access_token;
			//@ts-ignore
			const refresh = result.data.refresh_token;
			setAccess(access);
			setRefresh(refresh);
			console.log(result.data);
			userStore.setUser({
				username: result.data.user.username,
				email: result.data.user.email,
				avatar: "/avatars/user.jpg",
			});

			await goto("/");
		},
	});

	const { form: formData, enhance } = form;
</script>

<form class="p-6 md:p-8" method="POST" use:enhance>
	<div class="flex flex-col gap-6">
		<div class="flex flex-col items-center text-center">
			<h1 class="text-2xl font-bold">Welcome back</h1>
			<p class="text-muted-foreground text-balance">
				Login to yout TaskHub Account
			</p>
		</div>

		<Form.Field {form} name="usernameOrEmail" class="grid gap-3">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>Username</Form.Label>
					<Input
						{...props}
						bind:value={$formData.usernameOrEmail}
						type="text"
						placeholder="John Snow"
					/>
				{/snippet}
			</Form.Control>
			<Form.Description>
				You can either use your email or your username
			</Form.Description>
			<Form.FieldErrors />
		</Form.Field>

		<Form.Field {form} name="password" class="grid gap-3">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>Password</Form.Label>
					<Input {...props} bind:value={$formData.password} type="password" />
				{/snippet}
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>

		<Form.Button class="w-full">Login</Form.Button>
	</div>
</form>
