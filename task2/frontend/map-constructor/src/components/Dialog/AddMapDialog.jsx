import {Button, Dialog, DialogActions, DialogContent, DialogTitle} from "@mui/material";
import {TextInput} from "../../forms/TextInput";

export const AddMapDialog = ({
    open,
    onClose,
    setMapName,
    onSubmit
}) => {
    return (
        <Dialog open={open} onClose={onClose}>
            <DialogTitle>Добавить карту</DialogTitle>
            <DialogContent>
                <TextInput
                    label="Название"
                    setFieldValue={setMapName}
                >
                </TextInput>
            </DialogContent>
            <DialogActions>
                <Button onClick={onClose}>Отмена</Button>
                <Button onClick={onSubmit} variant="contained">Добавить</Button>
            </DialogActions>
        </Dialog>
    )
}