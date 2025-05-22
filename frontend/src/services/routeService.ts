import api from "../api";

export async function getRoutes(address: string, distance: string) {
  try {
    await api.post("/login")
    const response = await api.post("/location", { address, distance })

    return {response: response.data}
  } catch (err) {
    return {response: {}}
  }
}

export async function getPOIs(address: string) {
  try {
    await api.post("/login")
    const response = await api.post("/location/overpass", { address })

    return response.data
  } catch (err) {
  }
}
