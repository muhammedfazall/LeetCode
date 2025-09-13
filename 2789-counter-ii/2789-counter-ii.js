/**
 * @param {integer} init
 * @return { increment: Function, decrement: Function, reset: Function }
 */
var createCounter = function(init) {
    let curr = init;
    let obj={
        increment : function(){
           curr++;
           return curr;
        },
        decrement : function(){
           curr--;
           return curr;
        },
        reset : function(){
           curr = init;
           return curr;
        }
    }
    return obj;
};

/**
 * const counter = createCounter(5)
 * counter.increment(); // 6
 * counter.reset(); // 5
 * counter.decrement(); // 4
 */