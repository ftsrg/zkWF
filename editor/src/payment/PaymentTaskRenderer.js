import BaseRenderer from 'diagram-js/lib/draw/BaseRenderer';
import { append as svgAppend, create as svgCreate ,  attr as svgAttr,  remove as svgRemove} from 'tiny-svg';

const TASK_BORDER_RADIUS = 10;

class PaymentTaskRenderer extends BaseRenderer {
  constructor(eventBus, bpmnRenderer) {
    super(eventBus, 1500); // priority must be higher than default

    this.bpmnRenderer = bpmnRenderer;
  }

  canRender(element) {
    // Customize to only render the payment task type
    return element.type === 'bpmn:Task' && element.businessObject.$attrs.type === 'paymentTask';
  }

  drawShape(parentNode, element) {
    // Use the default shape for tasks as the base
    const taskShape = this.bpmnRenderer.drawShape(parentNode, element);

    const rect = drawRect(parentNode, 100, 80, TASK_BORDER_RADIUS, '#52B415');

    prependTo(rect, parentNode);

    svgRemove(taskShape);



    const dollarSign = svgCreate('text');
    dollarSign.textContent = '💵';

    svgAttr(dollarSign, {
      x: 5,
      y: 15,
      fill: 'green'
    });

    svgAppend(parentNode, dollarSign);

    // Centered text for participant name (static placeholder for now)
    const centerText = svgCreate('text');
    const participantName = element.businessObject.participant || '[Participant]';
    centerText.textContent = 'Pay ' + participantName;
    //centerText.setAttribute('x', '50%');
    //centerText.setAttribute('y', '50%');
    centerText.setAttribute('text-anchor', 'middle');
    centerText.setAttribute('dominant-baseline', 'middle');

    svgAttr(centerText, {
      transform: 'translate(50, 40)'
    });

    svgAppend(parentNode, centerText);

    // Bottom-right text for amount (static placeholder for now)
    const amountText = svgCreate('text');
    const amount = element.businessObject.amount || 0;
    amountText.textContent =  amount + ' ETH';
    //amountText.setAttribute('x', '80%'); // Adjust x/y for positioning
    //amountText.setAttribute('y', '90%');
    amountText.setAttribute('text-anchor', 'end');
    amountText.setAttribute('dominant-baseline', 'baseline');

    svgAttr(amountText, {
      transform: 'translate(95, 70)'
    });

    svgAppend(parentNode, amountText);

    return taskShape;
  }
}

PaymentTaskRenderer.$inject = ['eventBus', 'bpmnRenderer'];

export default PaymentTaskRenderer;

function drawRect(parentNode, width, height, borderRadius, strokeColor) {
  const rect = svgCreate('rect');

  svgAttr(rect, {
    width: width,
    height: height,
    rx: borderRadius,
    ry: borderRadius,
    stroke: strokeColor || '#000',
    strokeWidth: 2,
    fill: '#fff'
  });

  svgAppend(parentNode, rect);

  return rect;
}

// copied from https://github.com/bpmn-io/diagram-js/blob/master/lib/core/GraphicsFactory.js
function prependTo(newNode, parentNode, siblingNode) {
  parentNode.insertBefore(newNode, siblingNode || parentNode.firstChild);
}
