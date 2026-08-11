import { Events } from "@wailsio/runtime";
import { writable, get } from "svelte/store";
import { GetServerPlayers } from "../../bindings/seegolauncher/internal/services/api";

export let serverPlayers = writable<number>(0);
export let serverSlots = writable<number>(0);
export let serverAdmins = writable<number>(0);
export let serverQueue = writable<number>(0);
export let serverOnline = writable<boolean>(false)
export let queueTime = writable<number>(0)


export async function initServerStatus() {
  let server;
  try {
    server = await GetServerPlayers();
  } catch (err) {
    serverOnline.set(false)
    return;
  }

  if (server.players == 0) {
    serverOnline.set(false)
    return
  }

  serverPlayers.set(server.players)
  serverSlots.set(server.slots)
  serverAdmins.set(server.admins)
  serverQueue.set(server.queue)

  // without prio
  // about 1-3 players enters from the queue to the server in every 1 mins
  // update: i connected at 14:00 with 200 queue. After 2 hours, i was 150. My pos always changed between 140 and 160, so i will incrase the queue
  if (serverPlayers >= serverSlots) {
    // about 30-35 seconds one player disconnects // update up (changed to 50)
    // sometimes with the gold prio theres 100 queue, so without pro its like 30-35 seconds disconnect * 10
    if (get(serverQueue) > 0) {
      let totalSeconds = get(serverQueue) * 50;

      if (get(serverQueue) >= 100) {
        // queue * prio ratio 35% (+-5-20%) and + 120 seconds extra
        // its quite close, because one member in discord waited 7 hours with 500 queue without prio (connected in the morning, like 13-14:00, and connected at night)
        totalSeconds = totalSeconds + Math.floor(get(serverQueue) * 0.35) * 120
      }
      queueTime.set(totalSeconds);
    }
  } else {
    if (get(serverQueue) > 0) {
      // about 2 seconds one player connects
      const totalSeconds = get(serverQueue) * 2;
      queueTime.set(totalSeconds);
    } else {
      queueTime.set(0)
    }
  }

  serverOnline.set(true)
}
