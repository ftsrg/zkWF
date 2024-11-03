import { TextFieldEntry, isTextFieldEntryEdited } from '@bpmn-io/properties-panel';
import { useService } from 'bpmn-js-properties-panel'
import { html } from 'htm/preact';

export default function(element) {

  return [
    {
      id: 'publicKey',
      component: PublicKey,
      isEdited: isTextFieldEntryEdited
    }
  ];
}

function PublicKey(props) {
  const { element, id } = props;

  const modeling = useService('modeling');
  const translate = useService('translate');
  const debounce = useService('debounceInput');

  const getValue = () => {
    return element.businessObject.publicKey || '';
  }

  const setValue = value => {
    return modeling.updateProperties(element, {
      publicKey: value
    });
  }

  return html`<${TextFieldEntry}
    id=${ id }
    element=${ element }
    description=${ translate('The persons public key that will be used for this task') }
    label=${ translate('PublicKey') }
    getValue=${ getValue }
    setValue=${ setValue }
    debounce=${ debounce }
  />`;
}
