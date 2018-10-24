var App = function($dom) {
		this.$dom = $dom;
		this._initialize();
	};

var _ = App.prototype;

_._initialize = function() {
	this._fetchProblem();
	this._bindEvents();
};

_._fetchProblem = function() {
	var that = this;

	$.get('http://localhost:3000/api/fetchProblem', function(data) {
		that.problemStubs = data.problems.map(function(p) {
			return new ProblemStub(p);
		});
	});
};

_._bindEvents = function() {
	var that = this;

	this.$dom.find('.submit').click(function() {
		$.post('http://localhost:3000/api/submit', {
			input: JSON.stringify(that.problemStubs.map(function(ps) {
				return {
					id: ps.data.id,
					answer: ps.answerObjUser.getAnswer()
				};
			}))
		}, function(data) {
			that._renderResult(data.results);
			that.$dom.find('.submit').attr('disabled', true);
		});
	});
};

_._renderResult = function(data) {
	var that = this;

	this.problemStubs.map(function(ps, idx) {
		ps.fillAnswer(data[idx]);
	});

	data.forEach(function(d) {
		var $div = $('<div />');

		$div.html($('.tmpl-result-stub').html());
		$div.find('.id').text(d.id);
		if(d.result) {
			$div.find('.result').text('O');
		} else {
			$div.find('.result').text('X');
		}
		$div.appendTo('.main-result');
	});
};


var ProblemStub = function(data) {
		this.data = data;
		this._initialize();
	};

_ = ProblemStub.prototype;

_._initialize = function() {
	this._setAnswerObj();
	this._setDom();
};

_._setAnswerObj = function() {
	switch(this.data.type) {
		case 1:
			this.answerObjUser = new AnswerStubChoice(JSON.parse(this.data.choices));
			this.answerObjRight = new AnswerStubChoice(JSON.parse(this.data.choices));
			break;
		case 2:
			this.answerObjUser = new AnswerStubText();
			this.answerObjRight = new AnswerStubText();
			break;
		case 3:
			this.answerObjUser = new AnswerStubDraw();
			this.answerObjRight = new AnswerStubDraw();
			break;
		default:
			return;
	}
};

_._setDom = function() {
	this.$dom = $('<div class="problem-stub" />');
	this.$dom.html($('.tmpl-problem-stub').html());
	this.$dom.find('.id').text(this.data.id);
	this.$dom.find('.problem-text').text(this.data.problem_text);
	this.$dom.find('.answer-user').append(this.answerObjUser.$dom);

	this.$dom.appendTo('.problem');
};

_.fillAnswer = function(result) {
	this.answerObjUser.setLabel('나의 답: ');
	this.answerObjRight.setLabel('정답: ');

	this.answerObjRight.fillAnswer(result.answer);
	this.$dom.find('.answer-right').append(this.answerObjRight.$dom);
};


var AnswerStubChoice = function(choices) {
		this.choices = choices;
		this._initialize();
	};

_ = AnswerStubChoice.prototype;

_._initialize = function() {
	this._setDom();
};

_._setDom = function() {
	var that = this,
		rand = Math.floor(Math.random() * 1000000);

	this.$dom = $('<div class="answer-stub-choice" />');
	this.$dom.html($('.tmpl-answer-stub-choice').html());

	this.choices.forEach(function(c, idx) {
		that.$dom.find('input[type="radio"]').attr('name', 'choice_' + rand);
		that.$dom.find('.choice').eq(idx).text(c);
	});
};

_.getAnswer = function() {
	return this.$dom.find('input[type="radio"]:checked').val();
};

_.fillAnswer = function(value) {
	this.$dom.find('input[type="radio"][value="' + value + '"]').attr('checked', true);
};

_.setLabel = function(str) {
	this.$dom.find('.label').text(str);
};


var AnswerStubText = function() {
		this._initialize();
	};

_ = AnswerStubText.prototype;

_._initialize = function() {
	this._setDom();
};

_._setDom = function() {
	var that = this;

	this.$dom = $('<div class="answer-stub-text" />');
	this.$dom.html($('.tmpl-answer-stub-text').html());
};

_.getAnswer = function() {
	return this.$dom.find('.value').val();
};

_.fillAnswer = function(value) {
	this.$dom.find('.value').val(value);
};

_.setLabel = function(str) {
	this.$dom.find('.label').text(str);
};


var AnswerStubDraw = function() {
		this._initialize();
	};

_ = AnswerStubDraw.prototype;

_._initialize = function() {
	this._setDom();
	this._bindEvents();
};

_._setDom = function() {
	var that = this;

	this.$dom = $('<div class="answer-stub-draw" />');
	this.$dom.html($('.tmpl-answer-stub-draw').html());
	this.$dom.find('.canvas').attr({
		width: 125,
		height: 125
	});
};

_._bindEvents = function() {
	var that = this;

	this.ctx = this.$dom.find('.canvas')[0].getContext('2d');
	this.mousedown = false;

	this.$dom.find('.canvas').mousedown(function() {
		that.mousedown = true;
	});

	this.$dom.find('.canvas').mouseup(function() {
		that.mousedown = false;
	});

	this.$dom.find('.canvas').mouseleave(function() {
		that.mousedown = false;
	});

	this.$dom.find('.canvas').mousemove(function(e) {
		if(that.mousedown) {
			that.ctx.fillRect(e.offsetX, e.offsetY, 1, 1);
		}
	});
};

_.getAnswer = function() {
	return this.$dom.find('.canvas')[0].toDataURL();
};

_.fillAnswer = function(value) {
	var img = new Image();
	img.src = value;
	this.ctx.drawImage(img, 0, 0);
};

_.setLabel = function(str) {
	this.$dom.find('.label').text(str);
};
