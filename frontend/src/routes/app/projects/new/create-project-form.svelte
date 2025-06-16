<script lang="ts">
	import {
		type SuperValidated,
		type Infer,
		superForm,
	} from "sveltekit-superforms";
	import { zodClient } from "sveltekit-superforms/adapters";
	import * as Form from "$lib/components/ui/form";
	import { Input } from "$lib/components/ui/input";
	import { projectSchema, type ProjectSchema } from "./schema";

	let {
		projectForm,
	}: {
		projectForm: SuperValidated<Infer<ProjectSchema>>;
	} = $props();

	const form = superForm(projectForm, {
		validators: zodClient(projectSchema),
		onResult: async ({ result }) => {
			console.log(result);
		},
	});

	const { form: formData, enhance } = form;
</script>

<form class="p-6 md:p-8" method="POST" use:enhance>
	<div class="flex flex-col gap-6">
		<div class="flex flex-col items-center text-center">
			<h1 class="text-2xl font-bold">New Project</h1>
			<p class="text-muted-foreground text-balance">Create a new Project</p>
		</div>
	</div>
	<input type="hidden" name="formType" value="createTask" />
	<Form.Field {form} name="name" class="grid gap-3">
		<Form.Control>
			{#snippet children({ props })}
				<Form.Label>Name</Form.Label>
				<Input {...props} bind:value={$formData.name} type="text" />
			{/snippet}
		</Form.Control>
		<Form.FieldErrors />
	</Form.Field>

	<Form.Field {form} name="description" class="grid gap-3">
		<Form.Control>
			{#snippet children({ props })}
				<Form.Label>Description</Form.Label>
				<Input {...props} bind:value={$formData.description} type="text" />
			{/snippet}
		</Form.Control>
		<Form.FieldErrors />
	</Form.Field>

	<Form.Button
		class="w-full"
		onclick={() => {
			form.submit();
		}}>Create</Form.Button
	>
</form>
