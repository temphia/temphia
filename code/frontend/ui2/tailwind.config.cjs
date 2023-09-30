import { skeleton } from '@skeletonlabs/tw-plugin';
import { join } from 'path';


/** @type {import('tailwindcss').Config} */
module.exports = {
	darkMode: 'class',
	content: [
		'./src/**/*.{html,js,svelte,ts}',
		// 3. Append the path to the Skeleton package
		join(require.resolve(
			'@skeletonlabs/skeleton'),
			'../**/*.{html,js,svelte,ts}'
		)
	],
	theme: {
		extend: {},
	},
	plugins: [
		// 4. Append the Skeleton plugin (after other plugins)
		skeleton({
			themes: {
				// Register each theme within this array:
				preset: [ "skeleton", "modern", "crimson", "rocket" ] 
			}
		})
	]
}