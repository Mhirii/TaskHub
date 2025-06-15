<script lang="ts">
	import { Card, CardContent } from "$lib/components/ui/card";
	import * as Form from "$lib/components/ui/form";
	import { Input } from "$lib/components/ui/input";
	import * as Select from "$lib/components/ui/select/index.js";
	import { Textarea } from "$lib/components/ui/textarea/index.js";

	import { Calendar } from "$lib/components/ui/calendar/index.js";
	import * as Popover from "$lib/components/ui/popover/index.js";
	import CalendarIcon from "@lucide/svelte/icons/calendar";
	import { cn } from "$lib/utils.js";
	import { buttonVariants } from "$lib/components/ui/button/index.js";
	import {
		DateFormatter,
		type DateValue,
		getLocalTimeZone,
		today,
	} from "@internationalized/date";
	import {
		superForm,
		type SuperValidated,
		type Infer,
	} from "sveltekit-superforms";
	import { zodClient } from "sveltekit-superforms/adapters";
	import { taskSchema, type TaskUpdateSchema } from "./schema";
	import type { Task } from "$lib/types";
	import { ScrollArea } from "$lib/components/ui/scroll-area/index.js";

	const df = new DateFormatter("en-US", { dateStyle: "long" });

	export let data: { taskUpdateForm: SuperValidated<Infer<TaskUpdateSchema>> };
	export let task: Task;
	let buttonDisabled = true;
	let changeCouter = 0; // each time one of the form fields loads, it counts as a change
	console.log(task);

	const form = superForm(data.taskUpdateForm, {
		validators: zodClient(taskSchema),
		dataType: "json",
		onResult: async ({ result }) => {
			console.log(result);
		},
		onChange: () => {
			changeCouter++;
			if (changeCouter > 5) {
				buttonDisabled = false;
			}
		},
	});

	const { form: formData, enhance } = form;

	// Initialize form data with task values
	$formData.id = task.id;
	$formData.name = task.name;
	$formData.description = task.description;
	$formData.project_id = task.project_id;
	$formData.board_id = task.board_id;
	$formData.priority = task.priority;
	$formData.due_date = task.due_date
		? new Date(task.due_date * 1000)
		: undefined;
	$formData.assignee_id = task.assignee_id;

	let calendarValue: DateValue | undefined = task.due_date
		? today("UTC").set({
				year: new Date(task.due_date * 1000).getFullYear(),
				month: new Date(task.due_date * 1000).getMonth() + 1,
				day: new Date(task.due_date * 1000).getDate(),
			})
		: undefined;

	const priorities = ["None", "Low", "Medium", "High", "Urgent", "Critical"];
	const priorityOptions = [
		{ value: 0, label: priorities[0] },
		{ value: 1, label: priorities[1] },
		{ value: 2, label: priorities[2] },
		{ value: 3, label: priorities[3] },
		{ value: 4, label: priorities[4] },
		{ value: 5, label: priorities[5] },
	];
</script>

<Card class="w-full max-w-md lg:max-w-2xl shadow-md ">
	<ScrollArea class="h-full max-h-[70vh] w-full rounded-md ">
		<CardContent class="p-6">
			<form
				method="POST"
				use:enhance
				class="space-y-4"
				on:submit|preventDefault
			>
				<input type="hidden" name="formType" value="updateTask" />
				<div class="grid gap-4">
					<Form.Field {form} name="name">
						<Form.Control>
							{#snippet children({ props })}
								<Form.Label class="">Name</Form.Label>
								<Input
									{...props}
									bind:value={$formData.name}
									type="text"
									class=""
								/>
							{/snippet}
						</Form.Control>
						<Form.FieldErrors />
					</Form.Field>

					<Form.Field {form} name="description">
						<Form.Control>
							{#snippet children({ props })}
								<Form.Label class="">Description</Form.Label>
								<Textarea
									{...props}
									bind:value={$formData.description}
									class=""
								/>
							{/snippet}
						</Form.Control>
						<Form.FieldErrors />
					</Form.Field>

					<div class="grid gap-4 md:grid-cols-2">
						<Form.Field {form} name="priority" class="w-full">
							<Form.Control>
								{#snippet children({ props })}
									<Form.Label class="">Priority</Form.Label>
									<Select.Root
										{...props}
										type="single"
										value={$formData.priority?.toString()}
										onValueChange={(v) => ($formData.priority = Number(v))}
									>
										<Select.Trigger class="w-full">
											<Select.Label>
												{$formData.priority
													? priorities[$formData.priority]
													: "Select priority"}
											</Select.Label>
										</Select.Trigger>
										<Select.Content>
											{#each priorityOptions as option}
												<Select.Item
													value={`${option.value}`}
													label={option.label}
												/>
											{/each}
										</Select.Content>
									</Select.Root>
								{/snippet}
							</Form.Control>
							<Form.FieldErrors />
						</Form.Field>

						<Form.Field {form} name="due_date">
							<Form.Control>
								{#snippet children({ props })}
									<Form.Label class="">Due Date</Form.Label>
									<Popover.Root>
										<Popover.Trigger
											class={cn(
												buttonVariants({ variant: "outline" }),
												"justify-start text-left font-normal  w-full",
												!calendarValue && "text-muted-foreground",
											)}
										>
											<CalendarIcon class="mr-2 h-4 w-4" />
											{calendarValue
												? df.format(calendarValue.toDate(getLocalTimeZone()))
												: "Pick a date"}
										</Popover.Trigger>
										<Popover.Content class="w-auto p-0">
											<Calendar
												type="single"
												bind:value={calendarValue}
												minValue={today("UTC")}
												onValueChange={(v) => {
													$formData.due_date = v ? v.toDate("UTC") : undefined;
												}}
											/>
										</Popover.Content>
									</Popover.Root>
								{/snippet}
							</Form.Control>
							<Form.FieldErrors />
						</Form.Field>

						<Form.Field {form} name="assignee_id">
							<Form.Control>
								{#snippet children({ props })}
									<Form.Label class="">Assignee ID</Form.Label>
									<Input
										{...props}
										bind:value={$formData.assignee_id}
										type="number"
										class=""
									/>
								{/snippet}
							</Form.Control>
							<Form.FieldErrors />
						</Form.Field>

						<div class="grid gap-4 md:grid-cols-2">
							<div>
								<p class="text-sm font-medium">Created By</p>
								<p class="text-foreground/60 text-sm">User {task.created_by}</p>
							</div>
							<div>
								<p class="text-sm font-medium">Created</p>
								<p class="text-foreground/60 text-sm">
									{df.format(new Date(task.created_at * 1000))}
								</p>
							</div>
							<div>
								<p class="text-sm font-medium">Updated</p>
								<p class="text-foreground/60 text-sm">
									{df.format(new Date(task.updated_at * 1000))}
								</p>
							</div>
							<div>
								<p class="text-sm font-medium">Task ID</p>
								<p class="text-foreground/60 text-sm">{task.id}</p>
							</div>
						</div>

						<Form.Field {form} name="board_id" hidden>
							<Form.Control>
								{#snippet children({ props })}
									<Input
										{...props}
										bind:value={$formData.board_id}
										type="number"
									/>
								{/snippet}
							</Form.Control>
						</Form.Field>

						<Form.Field {form} name="project_id" hidden>
							<Form.Control>
								{#snippet children({ props })}
									<Input
										{...props}
										bind:value={$formData.project_id}
										type="number"
									/>
								{/snippet}
							</Form.Control>
						</Form.Field>
					</div>

					<div class="flex justify-end gap-2 mt-6">
						<Form.Button
							disabled={buttonDisabled}
							onclick={() => {
								console.log($formData);
							}}
						>
							Save
						</Form.Button>
					</div>
				</div>
			</form>
		</CardContent>
	</ScrollArea>
</Card>
