import {Box, Button, Dialog, DialogActions, DialogContent, DialogTitle} from "@mui/material";
import {TextInput} from "../../forms/TextInput";

export const AddCityDialog = ({
    open,
    onClose,
    setCityName,
    setY,
    setX,
    onSubmit
}) => {
    return (
        <Dialog open={open} onClose={onClose}>
            <DialogTitle>Добавить город</DialogTitle>
            <DialogContent>
                <Box sx={{
                    display: 'flex',
                    justifyContent: 'space-between',
                    alignItems: 'center',
                    flexDirection: 'column',
                }}>
                    <TextInput
                        label="Название"
                        setFieldValue={setCityName}
                    >
                    </TextInput>
                    <TextInput
                        label="Координата X"
                        setFieldValue={setX}
                    >
                    </TextInput>
                    <TextInput
                        label="Координата Y"
                        setFieldValue={setY}
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