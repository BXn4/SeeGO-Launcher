<script lang="ts">
    import logo from "../public/images/seego_nobg.svg";
</script>

<main>
    <div id="titlebar">
        <div id="drag-region">
            <div id="titlebar-hover-area"></div>
            <div id="window-title" class="text">
                <span>SeeGo Launcher</span>
            </div>
            <div id="window-controls">
                <div class="button" id="min-button" locale="minimize" title="">
                    <span class="icon">─</span>
                </div>
                <div class="button" id="close-button" locale="close" title="">
                    <span class="icon">✕</span>
                </div>
            </div>
        </div>
    </div>
</main>

<style>
    /* design from ggreborn app*/
    :root {
        --height: 40px;
        user-select: none;
    }

    #content {
        position: absolute;
        top: calc(var(--height));
        width: 100%;
        height: calc(100% - var(--height));
        z-index: -1;
    }

    #titlebar {
        background-color: #1f2736;
        display: block;
        position: fixed;
        height: var(--height);
        width: 100%;
        user-select: none;
        z-index: 3;
    }

    #titlebar #drag-region {
        --wails-draggable: drag;
        height: 100%;
        width: 100%;
        display: grid;
        grid-template-columns: auto 138px;
    }

    #titlebar-hover-area {
        position: absolute;
        display: block;
        position: fixed;
        height: 20px;
        width: 100%;
        user-select: none;
        z-index: 3;
        top: 30px;
        pointer-events: all;
    }

    #window-title {
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        font-family: AtkinsonHyperlegibleMono, "Courier New", Courier, monospace;
        font-size: 14px;
    }

    #window-controls {
        display: grid;
        grid-template-columns: repeat(2, 46px);
        position: absolute;
        top: 0;
        right: 0;
        height: 100%;
    }

    #window-controls .button {
        grid-row: 1 / span 1;
        display: flex;
        justify-content: center;
        align-items: center;

        transform: scale(1);
        transition: transform 0.08s ease-in-out;
        user-select: none;
    }

    #window-controls .button:hover {
        transform: scale(1);
        cursor: pointer;
    }

    #window-controls .button:active {
        transform: scale(1.2);
    }

    #min-button {
        grid-column: 1;
    }

    #close-button {
        grid-column: 2;
    }
</style>
