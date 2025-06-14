<script lang="ts">
	import { Button } from "$lib/components/ui/button";
	import type { Project } from "$lib/types";
	import { cn } from "$lib/utils";

	let {
		project,
		className,
		size,
	}: {
		project: Project;
		className?: string;
		size?: "sm" | "default";
	} = $props();
</script>

<Button class={cn("", className)} size={size || "default"}>
	<a href={`/app/projects/${project.id}`}>View Details</a>
</Button>
