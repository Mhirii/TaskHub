<script lang="ts">
	import { Button } from "$lib/components/ui/button";
	import { Plus } from "@lucide/svelte";
	import * as Dialog from "$lib/components/ui/dialog/index.js";
	import TaskForm from "./task-form.svelte";

	let { projectId, data, boardId } = $props();
</script>

<Dialog.Root>
	<Dialog.Trigger>
		<Button
			variant="outline"
			size="sm"
			class="h-8 w-8"
			onclick={() => console.log("add task")}
		>
			<Plus />
		</Button>
	</Dialog.Trigger>
	<Dialog.Content>
		<TaskForm {projectId} board={boardId} {data} />
	</Dialog.Content>
</Dialog.Root>
