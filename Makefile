vscode:
	deno run --allow-read --env-file --allow-net --allow-env supabase/functions/vscode/index.ts

chart:
	deno run --allow-read --env-file --allow-net --allow-env --watch public/index.ts

changelog:
	go run . changelog --version 2.127.0 --repo ../tokenscanner  