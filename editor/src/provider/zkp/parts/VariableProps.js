import { TextFieldEntry, isTextFieldEntryEdited } from '@bpmn-io/properties-panel';
import { useService } from 'bpmn-js-properties-panel'
import { html } from 'htm/preact';

export default function(element) {

  return [
    {
      id: 'variables',
      component: Variables,
      isEdited: isTextFieldEntryEdited
    }
  ];
}

function Variables(props) {
  const { element, id } = props;

  const modeling = useService('modeling');
  const translate = useService('translate');
  const debounce = useService('debounceInput');

  const getValue = () => {
    return element.businessObject.variables || '';
  }

  const setValue = value => {
    return modeling.updateProperties(element, {
      variables: value
    });
  }

  return html`<${TextFieldEntry}
    id=${ id }
    element=${ element }
    description=${ translate('The list of the global variables that the task can change') }
    label=${ translate('Variables') }
    getValue=${ getValue }
    setValue=${ setValue }
    debounce=${ debounce }
  />`;
}
