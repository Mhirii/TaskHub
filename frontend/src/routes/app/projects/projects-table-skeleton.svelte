<script lang="ts">
	import {
		Table,
		TableHeader,
		TableRow,
		TableHead,
		TableBody,
		TableCell,
	} from "$lib/components/ui/table";
	import { Skeleton } from "$lib/components/ui/skeleton";

	const skeletonRows = Array(5).fill(null);
</script>

<Table>
	<TableHeader>
		<TableRow>
			<TableHead><Skeleton class="h-4 w-20" /></TableHead>
			<TableHead><Skeleton class="h-4 w-32" /></TableHead>
			<TableHead><Skeleton class="h-4 w-20" /></TableHead>
			<TableHead><Skeleton class="h-4 w-16" /></TableHead>
			<TableHead><Skeleton class="h-4 w-24" /></TableHead>
			<TableHead class="hidden xl:table-cell"
				><Skeleton class="h-4 w-20" /></TableHead
			>
			<TableHead class="hidden xl:table-cell"
				><Skeleton class="h-4 w-20" /></TableHead
			>
			<TableHead><Skeleton class="h-4 w-16" /></TableHead>
		</TableRow>
	</TableHeader>
	<TableBody>
		{#each skeletonRows as _}
			<TableRow>
				<TableCell><Skeleton class="h-4 w-24" /></TableCell>
				<TableCell><Skeleton class="h-4 w-40" /></TableCell>
				<TableCell><Skeleton class="h-4 w-20" /></TableCell>
				<TableCell><Skeleton class="h-4 w-12" /></TableCell>
				<TableCell><Skeleton class="h-4 w-16" /></TableCell>
				<TableCell class="hidden xl:table-cell"
					><Skeleton class="h-4 w-20" /></TableCell
				>
				<TableCell class="hidden xl:table-cell"
					><Skeleton class="h-4 w-20" /></TableCell
				>
				<TableCell><Skeleton class="h-8 w-16" /></TableCell>
			</TableRow>
		{/each}
	</TableBody>
</Table>
