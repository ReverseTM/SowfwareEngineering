import {Box, Button, Dialog, DialogActions, DialogContent, DialogTitle} from "@mui/material";
import {TextInput} from "../../forms/TextInput";
import {SelectInput} from "../../forms/SelectInput";

export const UpdateCityDialog = ({
    open,
    onClose,
    cities,
    setOldCity,
    setNewCity,
    onSubmit
}) => {
    return (
        <Dialog open={open} onClose={onClose}>
            <DialogTitle>Изменить название города</DialogTitle>
            <DialogContent>
                <Box sx={{
                    display: 'flex',
                    flexDirection: 'column',
                    justifyContent: 'space-between',
                    alignItems: 'center',
                }}>
                    <SelectInput
                        label="Город"
                        items={cities}
                        setFieldValue={setOldCity}
                        sx={{
                            minWidth: '160px',
                            maxWidth: '160px'
                        }}
                    />
                    <TextInput
                        label="Новое название"
                        setFieldValue={setNewCity}
                    >
                    </TextInput>
                </Box>
            </DialogContent>
            <DialogActions>
                <Button onClick={onClose}>Отмена</Button>
                <Button onClick={onSubmit} variant="contained">Изменить</Button>
            </DialogActions>
        </Dialog>
    )
}