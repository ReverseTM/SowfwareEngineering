import {FormControl, Input, InputLabel} from "@mui/material";

export const TextInput = ({
    label,
    name,
    setFieldValue,
}) => {
    const onChange = (event) => {
        setFieldValue(event.target.value);
    }

    return (
        <FormControl variant="standard">
            <InputLabel>{label}</InputLabel>
            <Input
                name={name}
                onChange={onChange}
            />
        </FormControl>
    );
};