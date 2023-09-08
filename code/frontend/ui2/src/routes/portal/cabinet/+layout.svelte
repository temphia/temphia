<div class="h-full p-0 md:p-2">
    <div class="card p-2 variant-glass-primary h-full">
        <header class="card-header">
            <div class="flex justify-between">
                <ol class="breadcrumb">
                    <li class="crumb">
                        <a class="anchor" href="/elements/breadcrumbs">Home</a>
                    </li>
                    <li class="crumb-separator" aria-hidden>/</li>
                    <li class="crumb">
                        <a class="anchor" href="/elements/breadcrumbs">Bprints</a>
                    </li>
                </ol>

                <div>
                    <button type="button" class="btn variant-filled-primary">
                        <span>(icon)</span>
                        <span>Button</span>
                    </button>
                </div>
            </div>
        </header>
        <section class="p-2">
            <slot />
        </section>
    </div>
</div>
