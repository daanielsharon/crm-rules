import React, { useState } from 'react';
import { 
  Modal, 
  Box, 
  TextField, 
  Button, 
  Select, 
  MenuItem, 
  InputLabel, 
  FormControl,
  Typography
} from '@mui/material';

const SCHEDULE_OPTIONS = [
  { value: 'every_30_minutes', label: 'Every 30 Minutes' },
  { value: 'every_minute', label: 'Every Minute' },
  { value: 'hourly', label: 'Every Hour' },
  { value: 'every_5_minutes', label: 'Every 5 Minutes' },
  { value: 'every_10_minutes', label: 'Every 10 Minutes' },
];

const CONDITION_TEMPLATES = [
  { 
    label: 'Failed Logins', 
    template: 'failed_logins > {number}',
    placeholders: [{ 
      name: 'number', 
      type: 'number', 
      min: 0,
      helperText: 'Must be a non-negative number' 
    }]
  },
  { 
    label: 'Plan', 
    template: 'plan = {planType}',
    placeholders: [{ 
      name: 'planType', 
      type: 'select', 
      options: ['free', 'premium'] 
    }]
  }
];

const AddRuleModal = ({ open, onClose, onSubmit, rule, setRule }) => {
  const [conditionTemplate, setConditionTemplate] = useState(null);
  const [conditionValues, setConditionValues] = useState({});
  const [inputErrors, setInputErrors] = useState({});

  const handleConditionTemplateChange = (template) => {
    setConditionTemplate(template);
    setConditionValues({});
    setInputErrors({});
  };

  const handleConditionValueChange = (key, value) => {
    const placeholder = conditionTemplate.placeholders.find(p => p.name === key);
    
    // Validate input
    let isValid = true;
    if (placeholder.min !== undefined) {
      isValid = value >= placeholder.min;
    }

    // Update errors
    const newErrors = { 
      ...inputErrors, 
      [key]: !isValid 
    };
    setInputErrors(newErrors);

    if (isValid) {
      const newValues = { ...conditionValues, [key]: value };
      setConditionValues(newValues);

      // Generate full condition string
      if (conditionTemplate) {
        let fullCondition = conditionTemplate.template;
        Object.entries(newValues).forEach(([key, value]) => {
          // Add quotes for string values (select options)
          const formattedValue = placeholder.type === 'select' ? `'${value}'` : value;
          fullCondition = fullCondition.replace(`{${key}}`, formattedValue);
        });
        setRule({ ...rule, condition: fullCondition });
      }
    }
  };

  return (
    <Modal open={open} onClose={onClose}>
      <Box
        sx={{
          position: 'absolute',
          top: '50%',
          left: '50%',
          transform: 'translate(-50%, -50%)',
          width: 400,
          bgcolor: 'background.paper',
          p: 4,
          borderRadius: '8px',
        }}
      >
        <h2>Add Rule</h2>
        
        <TextField
          label="Rule Name"
          fullWidth
          margin="normal"
          value={rule.name || ''}
          onChange={(e) => setRule({ ...rule, name: e.target.value })}
        />

        <FormControl fullWidth margin="normal">
          <InputLabel>Condition Template</InputLabel>
          <Select
            value={conditionTemplate?.label || ''}
            onChange={(e) => {
              const selected = CONDITION_TEMPLATES.find(
                t => t.label === e.target.value
              );
              handleConditionTemplateChange(selected);
            }}
          >
            {CONDITION_TEMPLATES.map((template) => (
              <MenuItem key={template.label} value={template.label}>
                {template.label}
              </MenuItem>
            ))}
          </Select>
        </FormControl>

        {conditionTemplate && conditionTemplate.placeholders.map((placeholder) => (
          <TextField
            key={placeholder.name}
            label={placeholder.name}
            fullWidth
            margin="normal"
            type={placeholder.type}
            select={placeholder.type === 'select'}
            value={conditionValues[placeholder.name] || ''}
            onChange={(e) => handleConditionValueChange(
              placeholder.name, 
              e.target.value
            )}
            error={inputErrors[placeholder.name]}
            helperText={
              inputErrors[placeholder.name] 
                ? placeholder.helperText 
                : ''
            }
            inputProps={
              placeholder.min !== undefined 
                ? { min: placeholder.min } 
                : {}
            }
          >
            {placeholder.type === 'select' && 
              placeholder.options.map((option) => (
                <MenuItem key={option} value={option}>
                  {option}
                </MenuItem>
              ))
            }
          </TextField>
        ))}

        {rule.condition && (
          <Typography variant="body2" color="textSecondary" mt={2}>
            Condition: {rule.condition}
          </Typography>
        )}

        <FormControl fullWidth margin="normal">
          <InputLabel>Schedule</InputLabel>
          <Select
            value={rule.schedule || ''}
            onChange={(e) => setRule({ ...rule, schedule: e.target.value })}
          >
            {SCHEDULE_OPTIONS.map((option) => (
              <MenuItem key={option.value} value={option.value}>
                {option.label}
              </MenuItem>
            ))}
          </Select>
        </FormControl>

        <Box mt={2} display="flex" justifyContent="flex-end" gap={2}>
          <Button onClick={onClose} variant="outlined">
            Cancel
          </Button>
          <Button 
            onClick={onSubmit} 
            variant="contained" 
            color="primary"
            disabled={
              !rule.name || 
              !rule.condition || 
              !rule.schedule || 
              Object.values(inputErrors).some(error => error)
            }
          >
            Submit
          </Button>
        </Box>
      </Box>
    </Modal>
  );
};

export default AddRuleModal;
