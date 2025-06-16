<script lang="ts">
	import {
		type SuperValidated,
		type Infer,
		superForm,
	} from "sveltekit-superforms";
	import { zodClient } from "sveltekit-superforms/adapters";
	import * as Form from "$lib/components/ui/form";
	import { Input } from "$lib/components/ui/input";
	import { userSchema, type UserSchema } from "./schema";
	import { DateFormatter, type DateValue } from "@internationalized/date";
	import {
		Select,
		SelectContent,
		SelectGroup,
		SelectItem,
		SelectLabel,
		SelectTrigger,
	} from "$lib/components/ui/select";

	let {
		data,
	}: {
		data: { userForm: SuperValidated<Infer<UserSchema>> };
	} = $props();

	const form = superForm(data.userForm, {
		validators: zodClient(userSchema),
		onResult: async ({ result }) => {
			console.log(result);
		},
	});

	const { form: formData, enhance } = form;
</script>

<form class="p-6 md:p-8" method="POST" use:enhance>
	<div class="flex flex-col gap-6">
		<div class="flex flex-col items-center text-center">
			<h1 class="text-2xl font-bold">New User</h1>
			<p class="text-muted-foreground text-balance">Create a new user</p>
		</div>
	</div>
	<input type="hidden" name="formType" value="createUser" />
	<Form.Field {form} name="username" class="grid gap-3">
		<Form.Control>
			{#snippet children({ props })}
				<Form.Label>Username</Form.Label>
				<Input {...props} bind:value={$formData.username} type="text" />
			{/snippet}
		</Form.Control>
		<Form.FieldErrors />
	</Form.Field>

	<Form.Field {form} name="email" class="grid gap-3">
		<Form.Control>
			{#snippet children({ props })}
				<Form.Label>Email</Form.Label>
				<Input {...props} bind:value={$formData.email} type="email" />
			{/snippet}
		</Form.Control>
		<Form.FieldErrors />
	</Form.Field>

	<Form.Field {form} name="password" class="grid gap-3">
		<Form.Control>
			{#snippet children({ props })}
				<Form.Label>Password</Form.Label>
				<Input {...props} bind:value={$formData.password} type="text" />
			{/snippet}
		</Form.Control>
		<Form.FieldErrors />
	</Form.Field>

	<Form.Field {form} name="roles" class="grid gap-3">
		<Form.Control>
			{#snippet children({ props })}
				<Form.Label>Roles</Form.Label>
				<Select {...props} bind:value={$formData.roles}>
					<SelectTrigger>
						<SelectLabel>
							{$formData.roles ? $formData.roles.join(", ") : "Select a role"}
						</SelectLabel>
					</SelectTrigger>
					<SelectContent>
						<SelectGroup>
							<SelectItem value="admin">Admin</SelectItem>
							<SelectItem value="user">User</SelectItem>
							<SelectItem value="guest">Guest</SelectItem>
						</SelectGroup>
					</SelectContent>
				</Select>
			{/snippet}
		</Form.Control>
		<Form.FieldErrors />
		<Form.Description>
			Select the roles you want to assign to this user
		</Form.Description>
	</Form.Field>

	<Form.Button
		class="w-full"
		onclick={() => {
			form.submit();
		}}
	>
		Create
	</Form.Button>
</form>
