<script lang="ts">
    import { Events } from "@wailsio/runtime";
    import logo from "../public/images/seego_nobg.svg";
    import { Classes, Event, IDs } from "../utils/consts";
    import { getFromKey } from "../managers/localization";
    import { onMount } from "svelte";

    let setCurrrentProgressEvent: () => void;

    onMount(() => {
        setCurrrentProgressEvent = Events.On(
            Event.Splash.setCurrentProgress,
            async (e) => {
                const splashAlt = document.getElementById(IDs.Views.Splash.alt);
                splashAlt!.textContent = getFromKey(e.data);
            },
        );

        return () => {
            setCurrrentProgressEvent();
        };
    });
</script>

<main>
    <div id={IDs.Views.Splash.ID}>
        <img id={IDs.Views.Splash.logo} src={logo} alt="SeeGO Launcher logo" />
        <h1 class="{Classes.Global.text} {Classes.Views.Splash.title}">
            SeeGO Launcher
        </h1>
        <p
            id={IDs.Views.Splash.alt}
            class="{Classes.Global.comment} {Classes.Views.Splash.alt}"
        ></p>
        <span class={Classes.Views.Splash.loader}></span>
    </div>
</main>

<style>
    @import "../public/styles/views/splash.css";

    main {
        width: 100vw;
        height: 100vh;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    :global(body, html) {
        margin: 0;
        padding: 0;
    }
</style>
