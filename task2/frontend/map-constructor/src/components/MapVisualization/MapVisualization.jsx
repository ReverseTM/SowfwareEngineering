import { Box, Typography } from "@mui/material";
import React, { useRef, useEffect, useState } from "react";

export const MapVisualization = ({ cities, roads }) => {
    const scrollContainerRef = useRef(null);
    const [cellSize, setCellSize] = useState(40);
    const [gridDimensions, setGridDimensions] = useState({ width: 20, height: 20 });

    useEffect(() => {
        if (cities.length > 0) {
            const maxX = Math.max(...cities.map(c => c.x));
            const maxY = Math.max(...cities.map(c => c.y));
            setGridDimensions({
                width: Math.max(20, maxX + 2),
                height: Math.max(20, maxY + 2)
            });
        }
    }, [cities]);

    return (
        <Box sx={{
            width: `${cellSize * 21}px`,
            height: `${cellSize * 21}px`,
            overflow: 'auto',
            border: '1px solid #ccc'
        }} ref={scrollContainerRef}>
            <Box sx={{
                display: 'grid',
                gridTemplateColumns: `${cellSize}px repeat(${gridDimensions.width}, ${cellSize}px)`,
                gridTemplateRows: `${cellSize}px repeat(${gridDimensions.height}, ${cellSize}px)`,
                position: 'relative',
                minWidth: `${(gridDimensions.width + 1) * cellSize}px`,
                minHeight: `${(gridDimensions.height + 1) * cellSize}px`
            }}>
                {/* Пустая ячейка в левом верхнем углу */}
                <Box sx={{
                    gridColumn: 1,
                    gridRow: 1,
                    position: 'sticky',
                    top: 0,
                    left: 0,
                    zIndex: 5,
                    backgroundColor: 'white'
                }}></Box>

                {/* Нумерация столбцов */}
                {Array.from({ length: gridDimensions.width }).map((_, i) => (
                    <Box key={`col-${i}`} sx={{
                        gridColumn: i + 2,
                        gridRow: 1,
                        display: 'flex',
                        justifyContent: 'center',
                        alignItems: 'center',
                        position: 'sticky',
                        top: 0,
                        zIndex: 3,
                        backgroundColor: 'white'

                    }}>
                        <Typography variant="caption" fontWeight="bold">
                            {i + 1}
                        </Typography>
                    </Box>
                ))}

                {/* Нумерация строк */}
                {Array.from({ length: gridDimensions.height }).map((_, i) => (
                    <Box key={`row-${i}`} sx={{
                        gridColumn: 1,
                        gridRow: i + 2,
                        display: 'flex',
                        justifyContent: 'center',
                        alignItems: 'center',
                        position: 'sticky',
                        left: 0,
                        zIndex: 3,
                        backgroundColor: 'white'
                    }}>
                        <Typography variant="caption" fontWeight="bold">
                            {i + 1}
                        </Typography>
                    </Box>
                ))}

                {/* Основная сетка */}
                <Box
                    sx={{
                        gridColumn: `2 / ${gridDimensions.width + 2}`,
                        gridRow: `2 / ${gridDimensions.height + 2}`,
                        position: 'relative',
                        backgroundSize: `${cellSize}px ${cellSize}px`,
                        border: '1px solid #ccc',
                        backgroundColor: 'white',
                        backgroundImage: `
                            linear-gradient(to right, #eee 1px, transparent 1px),
                            linear-gradient(to bottom, #eee 1px, transparent 1px)
                        `,
                    }}
                >
                    {roads.map((road, index) => (
                        <Road key={`road-${index}`} road={road} cities={cities} cellSize={cellSize}/>
                    ))}

                    {cities.map(city => (
                        <City key={`city-${city.name}`} city={city} cellSize={cellSize}/>
                    ))}
                </Box>
            </Box>
        </Box>
    );
};

// Road и City остаются такими же, как у тебя

const Road = ({ road, cities, cellSize }) => {
    const fromCity = cities.find(c => c.name === road.from_city);
    const toCity = cities.find(c => c.name === road.to_city);

    if (!fromCity || !toCity) return null;

    const x1 = (fromCity.x - 1) * cellSize + cellSize / 2;
    const y1 = (fromCity.y - 1) * cellSize + cellSize / 2;
    const x2 = (toCity.x - 1) * cellSize + cellSize / 2;
    const y2 = (toCity.y - 1) * cellSize + cellSize / 2;
    const angle = Math.atan2(y2 - y1, x2 - x1) * 180 / Math.PI;

    return (
        <svg
            style={{
                position: 'absolute',
                top: 0,
                left: 0,
                width: '100%',
                height: '100%',
                pointerEvents: 'none',
                overflow: 'visible'
            }}
        >
            <defs>
                <marker id="arrowhead" markerWidth="10" markerHeight="7" refX="9" refY="3.5" orient="auto">
                    <polygon points="0 0, 10 3.5, 0 7" fill="#1976d2" />
                </marker>
            </defs>
            <line x1={x1} y1={y1} x2={x2} y2={y2} stroke="#1976d2" strokeWidth="2" />
            <text
                x={(x1 + x2) / 2}
                y={(y1 + y2) / 2}
                textAnchor="middle"
                fill="#d32f2f"
                fontSize={Math.min(12, cellSize / 3)}
                fontWeight="bold"
                transform={`rotate(${angle}, ${(x1 + x2) / 2}, ${(y1 + y2) / 2})`}
                dy="-5px"
            >
                {road.cost}
            </text>
        </svg>
    );
};

const City = ({ city, cellSize }) => {
    const fontSize = Math.min(12, cellSize / 4);

    return (
        <Box
            sx={{
                position: 'absolute',
                left: `${(city.x - 1) * cellSize}px`,
                top: `${(city.y - 1) * cellSize}px`,
                width: cellSize,
                height: cellSize,
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
                justifyContent: 'center',
                zIndex: 2
            }}
        >
            <Box
                sx={{
                    width: '100%',
                    textAlign: 'center',
                    whiteSpace: 'nowrap',
                    overflow: 'hidden',
                    textOverflow: 'ellipsis',
                    fontSize: fontSize,
                    fontWeight: 'bold',
                    mb: 0.5,
                    px: 0.5
                }}
            >
                {city.name}
            </Box>
            <Box
                sx={{
                    width: cellSize * 0.6,
                    height: cellSize * 0.6,
                    borderRadius: '50%',
                    backgroundColor: 'primary.main',
                    color: 'white',
                    display: 'flex',
                    alignItems: 'center',
                    justifyContent: 'center',
                    fontWeight: 'bold',
                    fontSize: Math.min(16, cellSize / 2)
                }}
            >
                {city.name.charAt(0).toUpperCase()}
            </Box>
        </Box>
    );
};
