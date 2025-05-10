const validateResponse = async response => {
    if (response.status !== 200 && response.status !== 201 && response.status !== 204) {
        await response.json()
            .then(response => response.error)
            .then(error => { throw new Error(error)} )
    }

    return response;
};

export const getAllMapNames = () =>
    fetch(`http://localhost:8080/api/maps/names`, {
        method: 'GET',
    })
        .then(response => validateResponse(response))
        .then(response => response.json())
        .then(response => response.data || []);

export const getAllCities = mapName =>
    fetch(`http://localhost:8080/api/maps/${mapName}/cities`, {
        method: 'GET',
    })
        .then(response => validateResponse(response))
        .then(response => response.json())
        .then(response => response.data || []);

export const getAllRoads = mapName =>
    fetch(`http://localhost:8080/api/maps/${mapName}/roads`, {
        method: 'GET',
    })
        .then(response => validateResponse(response))
        .then(response => response.json())
        .then(response => response.data || []);

export const addMap = mapName =>
    fetch(`http://localhost:8080/api/maps?name=${mapName}`, {
        method: 'POST',
    })
        .then(response => validateResponse(response))
        .then(response => response.json())
        .then(response => response.message)

export const addCity = (mapName, cityName, x, y) =>
    fetch(`http://localhost:8080/api/maps/${mapName}/cities`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            'city_name': cityName,
            'x': parseInt(x, 10),
            'y': parseInt(y, 10),
        }),
    })
        .then(response => validateResponse(response))
        .then(response => response.json())
        .then(response => response.message)
        .catch(err => { throw err; });

export const addRoad = (mapName, fromCity, toCity, cost) =>
    fetch(`http://localhost:8080/api/maps/${mapName}/roads`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            'from_city': fromCity,
            'to_city': toCity,
            'cost': parseInt(cost, 10)
        }),
    })
        .then(response => validateResponse(response))
        .then(response => response.json())
        .then(response => response.message)
        .catch(err => { throw err; });

export const updateCityName = (mapName, oldCityName, newCityName) =>
    fetch(`http://localhost:8080/api/maps/${mapName}/cities/${oldCityName}?name=${newCityName}`, {
        method: 'PATCH',
    })
        .then(response => validateResponse(response))
        .then(response => response.json())
        .then(response => response.message)

export const updateRoadCost = (mapName, fromCity, toCity, cost) =>
    fetch(`http://localhost:8080/api/maps/${mapName}/roads`, {
        method: 'PATCH',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            from_city: fromCity,
            to_city: toCity,
            cost: parseInt(cost, 10)
        })
    })
        .then(response => validateResponse(response))
        .then(response => response.json())
        .then(response => response.message)
        .catch(err => { throw err; });

export const deleteMap = mapName =>
    fetch(`http://localhost:8080/api/maps/${mapName}`, {
        method: 'DELETE',
    })
        .then(response => validateResponse(response))
        .then(response => response.json())
        .then(response => response.message)

export const deleteCity = (mapName, cityName) =>
    fetch(`http://localhost:8080/api/maps/${mapName}/cities/${cityName}`, {
        method: 'DELETE',
    })
        .then(response => validateResponse(response))
        .then(response => response.json())
        .then(response => response.message)

export const deleteRoad = (mapName, fromCity, toCity) =>
    fetch(`http://localhost:8080/api/maps/${mapName}/roads`, {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            from_city: fromCity,
            to_city: toCity,
        })
    })
        .then(response => validateResponse(response))
        .then(response => response.json())
        .then(response => response.message)

export const undo = mapName =>
    fetch(`http://localhost:8080/api/maps/${mapName}/undo`, {
        method: 'POST',
    })
        .then(response => validateResponse(response))

export const redo = mapName =>
    fetch(`http://localhost:8080/api/maps/${mapName}/redo`, {
        method: 'POST',
    })
        .then(response => validateResponse(response))

export const downloadMap = mapName =>
    fetch(`http://localhost:8080/api/maps/${mapName}/download`, {
        method: 'POST',
    })
        .then(response => validateResponse(response))
        .then(response => response.json())
        .then(response => response.data)

export const uploadMap = mapData =>
    fetch(`http://localhost:8080/api/maps/upload`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(mapData)
    })
        .then(response => validateResponse(response))
        .then(response => response.json())
        .then(response => response.message)