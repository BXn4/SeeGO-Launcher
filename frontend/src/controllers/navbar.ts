import { Events } from "@wailsio/runtime";
import {
  Event,
  currentNavbarActive,
} from "../utils/consts";

export async function SwitchNavTab(value: string) {
  currentNavbarActive.set(value);
  await Events.Emit(Event.Main.navigate, value);
}
