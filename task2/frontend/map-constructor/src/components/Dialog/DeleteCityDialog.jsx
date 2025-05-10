import {Button, Dialog, DialogActions, DialogContent, DialogTitle} from "@mui/material";
import {SelectInput} from "../../forms/SelectInput";

export const DeleteCityDialog = ({
    open,
    onClose,
    cities,
    setCityName,
    onSubmit
}) => {
    return (
        <Dialog open={open} onClose={onClose}>
            <DialogTitle>Удалить город</DialogTitle>
            <DialogContent>
                <SelectInput
                    label="Город"
                    items={cities}
                    setFieldValue={setCityName}
                    sx={{
                        minWidth: '160px',
                        maxWidth: '160px'
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