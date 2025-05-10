import {Box, Button, Dialog, DialogActions, DialogContent, DialogTitle} from "@mui/material";
import {TextInput} from "../../forms/TextInput";
import {SelectInput} from "../../forms/SelectInput";

export const AddRoadDialog = ({
    open,
    onClose,
    cities,
    setFromCity,
    setToCity,
    setCost,
    onSubmit
}) => {
    return (
        <Dialog open={open} onClose={onClose}>
            <DialogTitle>Добавить дорогу</DialogTitle>
            <DialogContent>
                <Box sx={{
                    display: "flex",
                    flexDirection: "column",
                    justifyContent: "space-between",
                    alignItems: "center",
                }}>
                    <SelectInput
                        label="Из города"
                        items={cities}
                        setFieldValue={setFromCity}
                        sx={{
                            minWidth: '160px',
                            maxWidth: '160px',
                        }}
                    />
                    <SelectInput
                        label="В город"
                        items={cities}
                        setFieldValue={setToCity}
                        sx={{
                            minWidth: '160px',
                            maxWidth: '160px',
                        }}
                    />
                    <TextInput
                        label="Стоимость"
                        setFieldValue={setCost}
                    >
                    </TextInput>
                </Box>
            </DialogContent>
            <DialogActions>
                <Button onClick={onClose}>Отмена</Button>
                <Button onClick={onSubmit} variant="contained">Добавить</Button>
            </DialogActions>
        </Dialog>
    )
}