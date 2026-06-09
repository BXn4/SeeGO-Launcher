<script lang="ts">
    import logo from "../public/images/seego_nobg.svg";
    import Titlebar from "./partials/Titlebar.svelte";
    import Navbar from "./partials/Navbar.svelte";
</script>

<main>
    <Titlebar></Titlebar>
    <Navbar></Navbar>
</main>

<style>
    main {
        width: 100vw;
        height: 100vh;
        overflow: hidden;
        position: relative;
        background: linear-gradient(
            -30deg,
            #1b2636,
            #1b2534,
            #1a2433,
            #1a2331,
            #1a222f,
            #1a212e,
            #19202c,
            #191f2a,
            #191e29,
            #181e27,
            #181d25,
            #171c24,
            #171b22,
            #161a21,
            #16191f
        );
    }
</style>
