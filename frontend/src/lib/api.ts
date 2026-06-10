import { GetCategories, GetItems } from "../../bindings/seegolauncher/internal/services/api"

export async function getCategories() {
  try {
    const categories = await GetCategories();
    return categories;
  } catch (err) {
    console.error("Failed to get categories:", err);
    throw err;
  }
}

export async function getItems(category: number) {
  try {
    const items = await GetItems(category);
    return items;
  } catch (err) {
    console.error("Failed to get items:", err);
    throw err;
  }
}
