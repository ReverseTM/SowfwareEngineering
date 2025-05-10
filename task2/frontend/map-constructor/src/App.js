import React, { useEffect, useState } from 'react';
import { MapVisualization } from './components/MapVisualization/MapVisualization';
import {
    Button,
    Box
} from '@mui/material';
import { SelectInput } from './forms/SelectInput';
import * as mapApi from './api/mapApi';
import {AddMapDialog} from "./components/Dialog/AddMapDialog";
import {DeleteMapDialog} from "./components/Dialog/DeleteMapDialog";
import {AddCityDialog} from "./components/Dialog/AddCityDialog";
import {UpdateCityDialog} from "./components/Dialog/UpdateCityDialog";
import {DeleteCityDialog} from "./components/Dialog/DeleteCityDialog";
import {AddRoadDialog} from "./components/Dialog/AddRoadDialog";
import {UpdateRoadDialog} from "./components/Dialog/UpdateRoadDialog";
import {DeleteRoadDialog} from "./components/Dialog/DeleteRoadDialog";

const App = () => {
    const [mapNames, setMapNames] = useState([]);
    const [selectedMap, setSelectedMap] = useState('');

    const [cities, setCities] = useState([]);
    const [cityNames, setCityNames] = useState([]);
    const [roads, setRoads] = useState([]);

    const [addMapOpen, setAddMapOpen] = useState(false);
    const [deleteMapOpen, setDeleteMapOpen] = useState(false);

    const [addCityOpen, setAddCityOpen] = useState(false);
    const [updateCityOpen, setUpdateCityOpen] = useState(false);
    const [deleteCityOpen, setDeleteCityOpen] = useState(false);

    const [addRoadOpen, setAddRoadOpen] = useState(false);
    const [updateRoadOpen, setUpdateRoadOpen] = useState(false);
    const [deleteRoadOpen, setDeleteRoadOpen] = useState(false);

    const [mapName, setMapName] = useState('');
    const [cityName, setCityName] = useState('');
    const [x, setX] = useState('');
    const [y, setY] = useState('');
    const [oldCityName, setOldCityName] = useState('');
    const [newCityName, setNewCityName] = useState('');
    const [fromCity, setFromCity] = useState('');
    const [toCity, setToCity] = useState('');
    const [cost, setCost] = useState('');

    useEffect(() => {
        loadMapNames(setMapNames);
    }, []);

    useEffect(() => {
        if (selectedMap) {
            loadCities(selectedMap, setCities, setCityNames);
            loadRoads(selectedMap, setRoads);
        } else {
            setCities([]);
            setCityNames([]);
            setRoads([]);
        }
    }, [selectedMap]);

    const handleAddMap = () => {
        mapApi.addMap(mapName)
            .then(message => console.log(message))
            .then(() => setAddMapOpen(false))
            .then(() => loadMapNames(setMapNames))
            .then(() => setSelectedMap(mapName))
            .catch(error => alert(error.message));
    }

    const handleDeleteMap = () => {
        mapApi.deleteMap(mapName)
            .then(message => console.log(message))
            .then(() => setDeleteMapOpen(false))
            .then(() => loadMapNames(setMapNames))
            .then(() => setSelectedMap(''))
            .catch(error => alert(error.message));
    }


    const handleAddCity = () => {
        mapApi.addCity(selectedMap, cityName, x, y)
            .then(message => console.log(message))
            .then(() => setAddCityOpen(false))
            .then(() => loadCities(selectedMap, setCities, setCityNames))
            .catch(error => alert(error.message));
    }

    const handleUpdateCityName = () => {
        mapApi.updateCityName(selectedMap, oldCityName, newCityName)
            .then(message => console.log(message))
            .then(() => setUpdateCityOpen(false))
            .then(() => loadCities(setMapNames, setCities, setCityNames))
            .catch(error => alert(error.message));
    }

    const handleDeleteCity = () => {
        mapApi.deleteCity(selectedMap, cityName)
            .then(message => console.log(message))
            .then(() => setDeleteCityOpen(false))
            .then(() => loadCities(selectedMap, setCities, setCityNames))
            .then(() => loadRoads(selectedMap, setRoads))
            .catch(error => alert(error.message));
    }

    const handleAddRoad = () => {
        mapApi.addRoad(selectedMap, fromCity, toCity, cost)
            .then(message => console.log(message))
            .then(() => setAddRoadOpen(false))
            .then(() => loadRoads(selectedMap, setRoads))
            .catch(error => alert(error.message));
    }

    const handleUpdateRoadCost = () => {
        mapApi.updateRoadCost(selectedMap, fromCity, toCity, cost)
            .then(message => console.log(message))
            .then(() => setUpdateRoadOpen(false))
            .then(() => loadRoads(selectedMap, setRoads))
            .catch(error => alert(error.message));
    }

    const handleDeleteRoad = () => {
        mapApi.deleteRoad(selectedMap, fromCity, toCity)
            .then(message => console.log(message))
            .then(() => setDeleteRoadOpen(false))
            .then(() => loadRoads(selectedMap, setRoads))
            .catch(error => alert(error.message));
    }

    const handleUndo = () => {
        mapApi.undo(selectedMap)
            .then(() => loadCities(selectedMap, setCities, setCityNames))
            .then(() => loadRoads(selectedMap, setRoads))
            .catch(error => alert(error.message));
    }

    const handleRedo = () => {
        mapApi.redo(selectedMap)
            .then(() => loadCities(selectedMap, setCities, setCityNames))
            .then(() => loadRoads(selectedMap, setRoads))
            .catch(error => alert(error.message));
    }

    const handleDownload = () => {
        mapApi.downloadMap(selectedMap)
            .then(data => {
                const blob = new Blob([[JSON.stringify(data, null, 2)]], { type: 'application/json' });
                const url = URL.createObjectURL(blob);

                const link = document.createElement('a');
                link.href = url;
                link.download = `${selectedMap}.json`;
                document.body.appendChild(link);
                link.click();
                document.body.removeChild(link);
                URL.revokeObjectURL(url);
            })
            .catch(error => alert(error.message));
    }

    const handleUpload = () => {
        const input = document.createElement('input');
        input.type = 'file';
        input.accept = 'application/json';

        input.onchange = () => {
            const file = input.files[0];
            if (!file) return;

            file.text()
                .then(text => {
                    const mapData = JSON.parse(text);
                    return mapApi.uploadMap(mapData);
                })
                .then(response => console.log(response))
                .then(() => loadMapNames(setMapNames))
                .catch(error => alert(error.message));
        };

        input.click();
    };

    return (
        <Box sx={{
            display: 'flex',
            flexDirection: 'row',
            alignItems: 'center',
            justifyContent: 'space-between',
            p: 3,
        }}>
            {/* Контейнер карты */}
            <Box sx={{
                flex: 1,
                width: '100%',
                height: '100%',
                display: 'flex',
                justifyContent: 'center',
                alignItems: 'center',
                p: 3,
            }}>
                <MapVisualization
                    cities={cities}
                    roads={roads}
                />
            </Box>

            {/* Панель управления */}
            <Box sx={{
                display: 'flex',
                flexDirection: 'row',
                width: '70%',
                gap: 2,
                p: 1,
                alignItems: 'center',
                justifyContent: 'space-between',
                zIndex: 1,
            }}>
                <Box sx={{
                    display: 'flex',
                    flexDirection: 'column',
                    gap: 2,
                }}>
                    <SelectInput
                        label="Карта"
                        setFieldValue={setSelectedMap}
                        items={mapNames}
                        value={selectedMap}
                        sx={{
                            maxWidth: '500px',
                            minWidth: '300px',
                            '& .MuiInputBase-root': {
                                fontSize: '0.875rem',
                            },
                        }}
                    />
                    <Button
                        variant="contained"
                        onClick={() => setAddMapOpen(true)}
                        sx={{ fontSize: '0.75rem' }}
                    >
                        Добавить карту
                    </Button>
                    <Button
                        variant="contained"
                        onClick={() => setDeleteMapOpen(true)}
                        sx={{ fontSize: '0.75rem' }}
                    >
                        Удалить карту
                    </Button>
                    <Button
                        variant="contained"
                        onClick={handleUpload}
                        sx={{ fontSize: '0.75rem' }}
                    >
                        Загрузить карту
                    </Button>
                    <Button
                        variant="contained"
                        onClick={handleDownload}
                        sx={{ fontSize: '0.75rem' }}
                    >
                        Выгрузить карту
                    </Button>
                    <Button
                        variant="contained"
                        onClick={() => setAddCityOpen(true)}
                        sx={{ fontSize: '0.75rem' }}
                    >
                        Добавить город
                    </Button>
                    <Button
                        variant="contained"
                        onClick={() => setUpdateCityOpen(true)}
                        sx={{ fontSize: '0.75rem' }}
                    >
                        Изменить имя города
                    </Button>
                    <Button
                        variant="contained"
                        onClick={() => setDeleteCityOpen(true)}
                        sx={{ fontSize: '0.75rem' }}
                    >
                        Удалить город
                    </Button>
                    <Button
                        variant="contained"
                        onClick={() => setAddRoadOpen(true)}
                        sx={{ fontSize: '0.75rem' }}
                    >
                        Добавить дорогу
                    </Button>
                    <Button
                        variant="contained"
                        onClick={() => setUpdateRoadOpen(true)}
                        sx={{ fontSize: '0.75rem' }}
                    >
                        Изменить стоимость дороги
                    </Button>
                    <Button
                        variant="contained"
                        onClick={() => setDeleteRoadOpen(true)}
                        sx={{ fontSize: '0.75rem' }}
                    >
                        Удалить дорогу
                    </Button>
                    <Button
                        variant="contained"
                        onClick={handleUndo}
                        sx={{ fontSize: '0.75rem' }}
                    >
                        Отменить действие
                    </Button>
                    <Button
                        variant="contained"
                        onClick={handleRedo}
                        sx={{ fontSize: '0.75rem' }}
                    >
                        Повторить действия
                    </Button>
                </Box>
            </Box>

            {/* Диалоговые окна */}
            <AddMapDialog
                open={addMapOpen}
                onClose={() => setAddMapOpen(false)}
                setMapName={setMapName}
                onSubmit={handleAddMap}
            />
            <DeleteMapDialog
                open={deleteMapOpen}
                onClose={() => setDeleteMapOpen(false)}
                mapNames={mapNames}
                setMapName={setMapName}
                onSubmit={handleDeleteMap}
            />
            <AddCityDialog
                open={addCityOpen}
                onClose={() => setAddCityOpen(false)}
                setCityName={setCityName}
                setX={setX}
                setY={setY}
                onSubmit={handleAddCity}
            />
            <UpdateCityDialog
                open={updateCityOpen}
                onClose={() => setUpdateCityOpen(false)}
                cities={cities}
                setOldCity={setOldCityName}
                setNewCity={setNewCityName}
                onSubmit={handleUpdateCityName}
            />
            <DeleteCityDialog
                open={deleteCityOpen}
                onClose={() => setDeleteCityOpen(false)}
                cities={cityNames}
                setCityName={setCityName}
                onSubmit={handleDeleteCity}
            />
            <AddRoadDialog
                open={addRoadOpen}
                onClose={() => setAddRoadOpen(false)}
                cities={cityNames}
                setFromCity={setFromCity}
                setToCity={setToCity}
                setCost={setCost}
                onSubmit={handleAddRoad}
            />
            <UpdateRoadDialog
                open={updateRoadOpen}
                onClose={() => setUpdateRoadOpen(false)}
                cities={cityNames}
                setFromCity={setFromCity}
                setToCity={setToCity}
                setCost={setCost}
                onSubmit={handleUpdateRoadCost}
            />
            <DeleteRoadDialog
                open={deleteRoadOpen}
                onClose={() => setDeleteRoadOpen(false)}
                cities={cityNames}
                setFromCity={setFromCity}
                setToCity={setToCity}
                onSubmit={handleDeleteRoad}
            />
        </Box>
    );
}

const loadMapNames = setMapNames => {
    mapApi
        .getAllMapNames()
        .then((mapNames => setMapNames(mapNames)))
        .catch(error => alert(error.message));
};

const loadCities = (mapName, setCities, setCityNames) => {
    mapApi
        .getAllCities(mapName)
        .then(cities => {
                setCities(cities)

                const cityNames = cities.map(city => city.name)
                setCityNames(cityNames)
            }
        )
        .catch(error => alert(error.message));
}

const loadRoads = (mapName, setRoads) => {
    mapApi.getAllRoads(mapName)
        .then(roads => setRoads(roads))
        .catch(error => alert(error.message));
}

export default App;