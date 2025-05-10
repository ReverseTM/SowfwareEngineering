import {Button, Dialog, DialogActions, DialogContent, DialogTitle} from "@mui/material";
import {SelectInput} from "../../forms/SelectInput";

export const DeleteMapDialog = ({
    open,
    onClose,
    mapNames,
    setMapName,
    onSubmit
}) => {
    return (
        <Dialog open={open} onClose={onClose}>
            <DialogTitle>Удалить карту</DialogTitle>
            <DialogContent>
                <SelectInput
                    label="Карта"
                    items={mapNames}
                    setFieldValue={setMapName}
                    sx = {{
                        minWidth: '140px',
                        maxWidth: '140px',
                    }}
                />
            </DialogContent>
            <DialogActions>
                <Button onClick={onClose}>Отмена</Button>
                <Button onClick={onSubmit} variant="contained">Удалить</Button>
            </DialogActions>
        </Dialog>
    )
}