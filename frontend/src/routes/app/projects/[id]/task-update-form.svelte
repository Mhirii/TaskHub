<script lang="ts">
	import {
		type SuperValidated,
		type Infer,
		superForm,
	} from "sveltekit-superforms";
	import { zodClient } from "sveltekit-superforms/adapters";
	import * as Form from "$lib/components/ui/form";
	import { Input } from "$lib/components/ui/input";
	import { taskSchema, type TaskUpdateSchema } from "./schema";
	import CalendarIcon from "@lucide/svelte/icons/calendar";
	import {
		DateFormatter,
		type DateValue,
		getLocalTimeZone,
	} from "@internationalized/date";
	import { cn } from "$lib/utils.js";
	import { buttonVariants } from "$lib/components/ui/button/index.js";
	import { Calendar } from "$lib/components/ui/calendar/index.js";
	import * as Popover from "$lib/components/ui/popover/index.js";
	import { today } from "@internationalized/date";
	import type { Task } from "$lib/types";

	const df = new DateFormatter("en-US", {
		dateStyle: "long",
	});

	let value = $state<DateValue | undefined>();
	let contentRef = $state<HTMLElement | null>(null);

	let {
		data,
		task,
	}: {
		data: { taskUpdateForm: SuperValidated<Infer<TaskUpdateSchema>> };
		task: Task;
	} = $props();

	const form = superForm(data.taskUpdateForm, {
		validators: zodClient(taskSchema),
		onResult: async ({ result }) => {
			console.log(result);
		},
	});

	const { form: formData, enhance } = form;
	$formData.project_id = task.project_id;
	$formData.board_id = task.board_id;
</script>

<form class="p-6 md:p-8" method="POST" use:form.enhance>
	<div class="flex flex-col gap-6">
		<div class="flex flex-col items-center text-center">
			<p class="text-xl font-semibold">{task.name}</p>
		</div>
	</div>
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

	<Form.Field {form} name="due_date" class="grid gap-3">
		<Form.Control>
			{#snippet children({ props })}
				<Form.Label>Due date</Form.Label>
				<Input {...props} bind:value={$formData.due_date} type="text" />

				<Popover.Root>
					<Popover.Trigger
						class={cn(
							buttonVariants({
								variant: "outline",
								class: " justify-start text-left font-normal",
							}),
							!value && "text-muted-foreground",
						)}
					>
						<CalendarIcon />
						{value
							? df.format(value.toDate(getLocalTimeZone()))
							: "Pick a date"}
					</Popover.Trigger>
					<Popover.Content bind:ref={contentRef} class="w-auto p-0">
						<Calendar
							type="single"
							bind:value
							minValue={today("UTC")}
							onValueChange={(v) => {
								if (v) {
									$formData.due_date = v.toDate("UTC");
								} else {
									$formData.due_date = undefined;
								}
							}}
						/>
					</Popover.Content>
				</Popover.Root>
			{/snippet}
		</Form.Control>
		<Form.FieldErrors />
	</Form.Field>

	<Form.Field {form} name="priority" class="grid gap-3">
		<Form.Control>
			{#snippet children({ props })}
				<Form.Label>Priority</Form.Label>
				<Input {...props} bind:value={$formData.priority} type="text" />
			{/snippet}
		</Form.Control>
		<Form.FieldErrors />
	</Form.Field>

	<Form.Field {form} name="assignee_id" class="grid gap-3">
		<Form.Control>
			{#snippet children({ props })}
				<Form.Label>Assignee</Form.Label>
				<Input {...props} bind:value={$formData.assignee_id} type="text" />
			{/snippet}
		</Form.Control>
		<Form.FieldErrors />
	</Form.Field>

	<Form.Field {form} name="board_id" class="grid gap-3" hidden>
		<Form.Control>
			{#snippet children({ props })}
				<Form.Label>Board</Form.Label>
				<Input {...props} bind:value={$formData.board_id} type="text" />
			{/snippet}
		</Form.Control>
		<Form.FieldErrors />
	</Form.Field>

	<Form.Field {form} name="project_id" class="grid gap-3" hidden>
		<Form.Control>
			{#snippet children({ props })}
				<Form.Label>Project</Form.Label>
				<Input {...props} bind:value={$formData.project_id} type="text" />
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
