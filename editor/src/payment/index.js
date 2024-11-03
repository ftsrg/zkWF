import PaymentTaskContextPad from './PaymentTaskContextPad';
import PaymentTaskPalette from './PaymentTaskPalette';
import PaymentTaskRenderer from './PaymentTaskRenderer';

export default {
  __init__: [ 'customContextPad', 'customPalette', 'customRenderer' ],
  customContextPad: [ 'type', PaymentTaskContextPad ],
  customPalette: [ 'type', PaymentTaskPalette ],
  customRenderer: [ 'type', PaymentTaskRenderer ]
};