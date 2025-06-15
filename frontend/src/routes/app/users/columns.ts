import type { ColumnDef } from "@tanstack/table-core";

export type User = {
	id: string;
	username: string;
	email: string;
	roles: string[];
	projects: Project[];
};

export const columns: ColumnDef<User>[] = [
	{
		accessorKey: "id",
		header: "ID",
	},
	{
		accessorKey: "email",
		header: "Email",
	},
	{
		accessorKey: "username",
		header: "Username",
	},
	{
		accessorKey: "roles",
		header: "Roles",
		cell: ({ cell }) => {
			return cell.getValue().join(", ");
		},
	},
	{
		accessorKey: "projects",
		header: "Projects",
		cell: ({ cell }) => {
			return cell.getValue()?.map((p) => p.name)?.join(", ") || "None";
		},
	}
];
