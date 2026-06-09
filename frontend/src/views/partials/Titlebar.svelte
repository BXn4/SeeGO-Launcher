<script lang="ts">
    import { Events } from "@wailsio/runtime";
    async function closeWindow() {
        await Events.Emit("close", null);
    }
    async function minimizeWindow() {
        await Events.Emit("minimize", null);
    }
    async function toggleMaximize() {
        await Events.Emit("toggle-maximize", null);
    }
</script>

<main>
    <div
        id="titlebar"
        role="button"
        tabindex="0"
        ondblclick={() => toggleMaximize()}
    >
        <div id="drag-region">
            <div id="titlebar-hover-area"></div>
            <div id="window-title">
                <span class="dim1">See</span>
                <span class="dim2">Go</span>
            </div>
            <div id="window-controls">
                <button
                    class="window-control"
                    id="min-button"
                    title="Minimize"
                    onclick={() => minimizeWindow()}
                >
                    <span>−</span>
                </button>
                <button
                    class="window-control"
                    id="close-button"
                    title="Close"
                    onclick={() => closeWindow()}
                >
                    <span>×</span>
                </button>
            </div>
        </div>
    </div>
</main>
