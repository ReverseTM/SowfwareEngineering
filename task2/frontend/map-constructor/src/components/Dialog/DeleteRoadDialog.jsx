import {Box, Button, Dialog, DialogActions, DialogContent, DialogTitle} from "@mui/material";
import {SelectInput} from "../../forms/SelectInput";

export const DeleteRoadDialog = ({
    open,
    onClose,
    cities,
    setFromCity,
    setToCity,
    onSubmit
}) => {
    return (
        <Dialog open={open} onClose={onClose}>
            <DialogTitle>Удалить дорогу</DialogTitle>
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
                </Box>
            </DialogContent>
            <DialogActions>
                <Button onClick={onClose}>Отмена</Button>
                <Button onClick={onSubmit} variant="contained">Удалить</Button>
            </DialogActions>
        </Dialog>
    )
}